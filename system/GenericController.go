package system

import (
	"action"
	"fmt"
	"os"
	"strings"
)

type GenericController struct {
	Id           string
	Namespace    string
	RelAddr      string
	AgeClustered bool
	Model        [4]*GenericModel
}

func (me *GenericController) NewGenericModel(op action.Action, fieldset []*Field) *GenericModel {
	var model = &GenericModel{
		Name:         me.Id,
		Suffix:       "",
		AgeClustered: me.AgeClustered,
		Fieldset:     fieldset,
	}

	if op == action.NONE {
		return model
	}

	if op == action.GET {
		model.Suffix = "List"
	}

	me.Model[op] = model

	return me.Model[op]
}

func (me *GenericController) getDatasetImpl() string {
	if me.AgeClustered {
		return "\n\t\t$request = new Request();\n\t\treturn _GET::find([\"conditions\" => \"_age = :_age:\", \"bind\" => [\"_age\" => $request->getHeader(\"Age\")]]);\n\t"
	} else {
		return "\n\t\treturn _GET::find();\n\t"
	}
}

func (me *GenericController) Publish(addr string) error {
	var byt []byte
	var e error
	if byt, e = os.ReadFile("system\\GenericController.php"); e != nil {
		return e
	}

	if me.Model[action.GET] == nil {
		me.Model[action.GET] = me.Model[action.SET]
	}

	if me.Model[action.SET] == nil {
		me.Model[action.SET] = me.Model[action.GET]
	}

	// var updbutton = "{ \\\"data\\\": null, \\\"className\\\": \\\"dt-center editor-edit\\\", \\\"defaultContent\\\": \\\"<i class=\\\\\\\"fa fa-pencil\\\\\\\" />\\\", \\\"orderable\\\": false }"
	// var rembutton = "{ \\\"data\\\": null, \\\"className\\\": \\\"dt-center editor-delete\\\", \\\"defaultContent\\\": \\\"<i class=\\\\\\\"fa fa-trash\\\\\\\" />\\\", \\\"orderable\\\": false }"
	var model = "use %s\\Models\\%s as Enum_%d;\n"
	var enum = "\n\t\t$enum_%d = Enum_%d::find([\"order\" => \"%s\"]);\n\t\t$this->view->setVar(\"enum_%d\", $enum_%d);\n\t"
	var getter = "\n\t\t$model->%s = $this->getViewStateValue(\"%s\");"

	var lib = ""
	var optionset = ""
	var modelreceiver = "\n\t\t$model = new _SET();\n\t\t$id = $this->getViewStateValue(\"_id\");\n\t\tif ($id != NULL) {\n\t\t\t$model = _SET::findFirst([\n\t\t\t\t\"conditions\" => \"_id = :_id:\",\n\t\t\t\t\"bind\" => [\"_id\" => $id]\n\t\t\t]);\n\t\t}\n"
	for i, field := range me.Model[action.SET].Fieldset {
		modelreceiver += fmt.Sprintf(getter, field.Name, field.Name)

		if field.Optionset != nil {
			lib += fmt.Sprintf(model, me.Namespace, field.Optionset.Name, i)
			optionset += fmt.Sprintf(enum, i, i, field.Optionset.Value, i, i)
		}
	}

	modelreceiver += "\n\t\treturn $model;"
	// rowformat = append(rowformat, updbutton)
	// rowformat = append(rowformat, rembutton)

	var php = string(byt)
	php = strings.Replace(string(php), "namespace System", fmt.Sprintf("namespace %s", me.Namespace), -1)
	php = strings.Replace(string(php), "use System\\Models\\GenericModel as _GET", fmt.Sprintf("use %s\\Models\\%s%s as _GET", me.Namespace, me.Model[action.GET].Name, me.Model[action.GET].Suffix), -1)
	php = strings.Replace(string(php), "use System\\Models\\GenericModel as _SET", fmt.Sprintf("use %s\\Models\\%s%s as _SET", me.Namespace, me.Model[action.SET].Name, me.Model[action.SET].Suffix), -1)
	php = strings.Replace(string(php), "use System\\Enums;", lib, -1)
	php = strings.Replace(string(php), "abstract class GenericController", fmt.Sprintf("class %sController", me.Id), -1)

	php = strings.Replace(string(php), "abstract function getViewState(): _SET;", fmt.Sprintf("private function getViewState(): _SET {%s\n\t}\n\t", modelreceiver), -1)
	php = strings.Replace(string(php), "abstract function getDataset(): Object;", fmt.Sprintf("private function getDataset(): Object {%s}\n\t", me.getDatasetImpl()), -1)
	php = strings.Replace(string(php), "abstract function dropdown(): void;", fmt.Sprintf("private function dropdown(): void {%s}", optionset), -1)

	var arr1 []string
	var arr2 []string
	if me.Model[action.SET].Suffix == "List" {
		arr1 = strings.Split(php, "\n\tpublic function newAction()")
		arr2 = strings.Split(arr1[1], "} // new")
		php = arr1[0] + arr2[1]

		arr1 = strings.Split(php, "\n\tpublic function getAction()")
		arr2 = strings.Split(arr1[1], "} // get")
		php = arr1[0] + arr2[1]

		arr1 = strings.Split(php, "\n\tpublic function putAction()")
		arr2 = strings.Split(arr1[1], "} // put")
		php = arr1[0] + arr2[1]

		arr1 = strings.Split(php, "\n\tpublic function remAction()")
		arr2 = strings.Split(arr1[1], "} // rem")
		php = arr1[0] + arr2[1]
	}

	byt = []byte(php)
	if e = os.WriteFile(fmt.Sprintf("%s\\controllers\\%sController.php", addr, me.Id), byt, 0644); e != nil {
		return e
	}

	var view *GenericView
	for i, model := range me.Model {
		var op = action.Action(i)
		if model == nil || (op == action.GET && model.Suffix == "") {
			continue
		}

		if e = model.Publish(me.Namespace, addr); e != nil {
			return e
		}

		if op == action.GET {
			continue
		}

		view = NewGenericView(me.Id, model.Fieldset)
		if e = view.Publish(addr, op); e != nil {
			return e
		}
	}

	return e
}
