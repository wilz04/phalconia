<?php
namespace System\Controllers;

use System\Models\GenericModel as _GET;
use System\Models\GenericModel as _SET;
use System\Enums;
use Phactory\Auth\SecurityUsers as Authorization;
use Phalcon\Http\Response;
use Phalcon\Http\Request;
use Phalcon\Mvc\Controller;

abstract class GenericController extends Controller {

	private Request $request;
	private Response $response;

	public function initialize() {
		// Create a request and response instance
		$this->request = new Request();
		$this->response = new Response();
	}

	abstract function getViewState(): _SET;
	abstract function getDataset(): Object;
	abstract function dropdown(): void;

	private function getListId(): string {
		return sprintf("_%s_list", strtolower(_GET::getName()));
	}

	private function getViewStateValue(string $key): ?string {
		$value = $this->request->getPost($key);
		return !empty($value)? $value: NULL;
	}

	private function getToolbarRights(): string {
		$ctrl = strtolower(_GET::getName());
		$from = $this->request->getHeader("From");
		return sprintf(
			"{\"new\": %b, \"put\": %b, \"rem\": %b}",
			is_callable([$this, "GETACTION"]) && Authorization::allow($ctrl, "new", $from),
			is_callable([$this, "PUTACTION"]) && Authorization::allow($ctrl, "put", $from),
			is_callable([$this, "REMACTION"]) && Authorization::allow($ctrl, "rem", $from)
		);
	}

	public function openAction(): Response {
		if (!Authorization::allow(strtolower(_GET::getName()), "open", $this->request->getHeader("From"))) {
			$this->response->setStatusCode(401, "Unauthorized");
			return $this->response;
		}

		$this->view->disable();

		$content = sprintf(
			"{ \"id\": \"%s\", \"src\": \"%s\", \"fieldset\": %s, \"rowformat\": %s, \"headset\": [\"%s\"], \"rights\": %s }",
			$this->getListId(),
			strtolower(_GET::getName()),
			(is_callable([$this, "PUTACTION"])? _SET::getSerializedFieldset(): _GET::getSerializedFieldset()),
			_GET::getRowFormat(),
			implode("\", \"", _GET::getHeadset()),
			$this->getToolbarRights(),
		);

		$this->response->setContent($content);
		$this->response->setContentType("application/json");
		return $this->response;
	}

	public function listAction(): Response {
		if (!Authorization::allow(strtolower(_GET::getName()), "list", $this->request->getHeader("From"))) {
			$this->response->setStatusCode(401, "Unauthorized");
			return $this->response;
		}

		$this->view->disable();

		$dataset = $this->getDataset();
		$headset = _GET::getHeadset();
		$body = [];
		foreach ($dataset as $data) {
			$row = [];
			$j = 0;
			foreach ($headset as $i => $head) {
				$row[] = sprintf("\"%s\": \"%s\"", $j++, str_replace("\"", "\\\"", $data->$i));
			}

			$body[] = sprintf("{ %s }", implode(", ", $row));
		}

		$content = sprintf("{ \"data\": [%s] }", implode(", ", $body));
		$this->response->setContent($content);
		$this->response->setContentType("application/json");
		return $this->response;
	}

	public function newAction(): Response {
		$this->dropdown();
		$this->response->setContent($this->view->getRender(_SET::getName(), "set"));
		return $this->response;
	} // new

	public function getAction(): Response {
		if (!Authorization::allow(strtolower(_SET::getName()), "get", $this->request->getHeader("From"))) {
			$this->response->setStatusCode(401, "Unauthorized");
			return $this->response;
		}

		$this->dropdown();
		$this->view->setVar("identity", "none");

		$model = _SET::findFirst([
			"conditions" => "_id = :_id:",
			"bind" => ["_id" => $this->request->getPost("_id")]
		]);

		$headset = _SET::getHeadset();
		foreach ($headset as $i => $head) {
			$this->view->setVar($i, $model->$i);
		}

		$this->response->setContent($this->view->getRender(_SET::getName(), "set"));
		return $this->response;
	} // get

	public function putAction(): Response {
		if (!Authorization::allow(strtolower(_SET::getName()), "put", $this->request->getHeader("From"))) {
			$this->response->setStatusCode(401, "Unauthorized");
			return $this->response;
		}

		$this->view->disable();

		$model = $this->getViewState();
		$ok = false;
		try {
			$ok = $model->save();
		} catch (\Exception $e) {
			header("Warning: ".$e->getMessage());
		}

		if ($ok === false) {
			$this->response->setStatusCode(409, "Conflict");
		} else {
			$this->response->setStatusCode(200, "OK");
		}

		$this->response->setContent("{}");
		$this->response->setContentType("application/json");
		return $this->response;
	} // put

	public function remAction(): Response {
		if (!Authorization::allow(strtolower(_SET::getName()), "rem", $this->request->getHeader("From"))) {
			$this->response->setStatusCode(401, "Unauthorized");
			return $this->response;
		}

		$this->view->disable();

		$model = _SET::find(sprintf("_id in ('%s')", implode("', '", $this->request->getPost("_id"))));

		$ok = false;
		try {
			$ok = $model->delete();
		} catch (\Exception $e) {
			header("Warning: ".$e->getMessage());
		}

		if ($ok === false) {
			$this->response->setStatusCode(405, "Method Not Allowed");
		} else {
			$this->response->setStatusCode(200, "OK");
		}

		$this->response->setContent("{}");
		$this->response->setContentType("application/json");
		return $this->response;
	} // rem

}
?>