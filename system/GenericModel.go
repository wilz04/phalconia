package system

import (
	"fmt"
	"os"
	"strings"
)

type GenericModel struct {
	Name         string
	Suffix       string
	AgeClustered bool
	Fieldset     []*Field
}

func (me *GenericModel) Publish(namespace string, addr string) error {
	var byt []byte
	var e error
	if byt, e = os.ReadFile("system\\GenericModel.php"); e != nil {
		return e
	}

	var prop = "[\"label\" => \"%s\", \"name\" => \"%s\", \"type\" => \"%s\"]"
	var fieldset []string
	for _, field := range me.Fieldset {
		fieldset = append(fieldset, fmt.Sprintf(prop, field.Description, field.Name, field.Ft))
	}

	var php = string(byt)
	php = strings.Replace(string(php), "namespace System", fmt.Sprintf("namespace %s", namespace), -1)
	php = strings.Replace(string(php), "abstract class GenericModel", fmt.Sprintf("class %s%s", me.Name, me.Suffix), -1)
	php = strings.Replace(string(php), "abstract function getFieldset(): Array;", fmt.Sprintf("public static function getFieldset(): Array {\n\t\treturn [%s];\n\t}\n", strings.Join(fieldset, ", ")), -1)

	byt = []byte(php)
	if e = os.WriteFile(fmt.Sprintf("%s\\models\\%s%s.php", addr, me.Name, me.Suffix), byt, 0644); e != nil {
		return e
	}

	return e
}
