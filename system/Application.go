package system

import (
	"fmt"
	"os"
	"strings"
)

type Application struct {
	Modules []*Module
}

func NewApplication() (me *Application) {
	me = &Application{}

	return
}

func (me *Application) NewModule(id string, namespace string, relAddr string) *Module {
	var module = &Module{
		Id:        id,
		Namespace: namespace,
		RelAddr:   fmt.Sprintf("%s/%s", relAddr, strings.ToLower(id)),
	}

	me.Modules = append(me.Modules, module)

	return module
}

func (me *Application) Publish(addr string) error {
	var e error
	var f = "function %s(%s): %s {\n%s}\n\n"
	var route = "\t$router->add(\"/%s/%s\", [\n"
	route += "\t\t\"module\" => \"%s\",\n"
	route += "\t\t\"controller\" => \"%s\",\n"
	route += "\t\t\"action\" => \"%s\",\n"
	route += "\t]);\n\n"
	var moduleregister = "\treturn [\n"
	moduleregister += "\t\t\"className\" => \\%s\\%s\\%s::class,\n"
	moduleregister += "\t\t\"path\" => \"%s/%s.php\",\n"
	moduleregister += "\t];\n"

	if e = os.Mkdir(addr+"\\system", os.ModePerm); e != nil {
		return e
	}

	var actions = ""
	var registers = ""
	var php = ""
	for _, module := range me.Modules {
		if e = module.Publish(addr + "\\system"); e != nil {
			return e
		}

		actions = ""
		for _, ctrl := range module.Controllers {
			actions += fmt.Sprintf(route, strings.ToLower(ctrl.Id), "open", strings.ToLower(module.Id), ctrl.Id, "open")
			actions += fmt.Sprintf(route, strings.ToLower(ctrl.Id), "list", strings.ToLower(module.Id), ctrl.Id, "list")
			actions += fmt.Sprintf(route, strings.ToLower(ctrl.Id), "new", strings.ToLower(module.Id), ctrl.Id, "new")
			actions += fmt.Sprintf(route, strings.ToLower(ctrl.Id), "get", strings.ToLower(module.Id), ctrl.Id, "get")
			actions += fmt.Sprintf(route, strings.ToLower(ctrl.Id), "put", strings.ToLower(module.Id), ctrl.Id, "put")
			actions += fmt.Sprintf(route, strings.ToLower(ctrl.Id), "rem", strings.ToLower(module.Id), ctrl.Id, "rem")
		}

		registers = fmt.Sprintf(moduleregister, module.Namespace, module.Id, module.Id, module.RelAddr, module.Id)

		php += fmt.Sprintf(f, fmt.Sprintf("set%sAction", module.Id), "Router $router", "void", actions)
		php += fmt.Sprintf(f, fmt.Sprintf("get%sRegister", module.Id), "", "array", registers)
	}

	php = fmt.Sprintf("<?php\n%s\n?>", php)
	var byt = []byte(php)
	if e = os.WriteFile(addr+"\\index.php", byt, 0644); e != nil {
		return e
	}

	return e
}
