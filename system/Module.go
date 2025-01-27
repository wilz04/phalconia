package system

import (
	"fmt"
	"os"
	"strings"
)

type Module struct {
	Id          string
	Namespace   string
	RelAddr     string
	Controllers []*GenericController
}

func (me *Module) NewGenericController(id string, ageclustered bool) *GenericController {
	var ctrl = &GenericController{
		Id:           me.Id + id,
		Namespace:    fmt.Sprintf("%s\\%s", me.Namespace, me.Id),
		RelAddr:      me.RelAddr,
		AgeClustered: ageclustered,
	}

	me.Controllers = append(me.Controllers, ctrl)

	return ctrl
}

func (me *Module) Publish(addr string) error {
	var byt []byte
	var e error
	if byt, e = os.ReadFile("system\\Module.php"); e != nil {
		return e
	}

	var php = string(byt)
	php = strings.Replace(string(php), "namespace system", fmt.Sprintf("namespace %s\\%s", me.Namespace, me.Id), -1)
	php = strings.Replace(string(php), "abstract class Module", fmt.Sprintf("class %s", me.Id), -1)

	php = strings.Replace(string(php), "abstract function getName(): string;", fmt.Sprintf("private function getName(): string {\n\t\treturn \"%s\\%s\";\n\t}\n", me.Namespace, me.Id), -1)
	php = strings.Replace(string(php), "abstract function getRelativeUrl(): string;", fmt.Sprintf("private function getRelativeUrl(): string {\n\t\treturn \"%s\";\n\t}", me.RelAddr), -1)

	addr += "\\" + strings.ToLower(me.Id)
	if e = os.Mkdir(addr, os.ModePerm); e != nil {
		return e
	}

	byt = []byte(php)
	if e = os.WriteFile(addr+"\\"+me.Id+".php", byt, 0644); e != nil {
		return e
	}

	if e = os.Mkdir(addr+"\\models", os.ModePerm); e != nil {
		return e
	}

	if e = os.Mkdir(addr+"\\views", os.ModePerm); e != nil {
		return e
	}

	if e = os.Mkdir(addr+"\\controllers", os.ModePerm); e != nil {
		return e
	}

	for _, ctrl := range me.Controllers {
		ctrl.Publish(addr)
	}

	return e
}
