package system

import (
	"action"
	"fmt"
	"os"
)

type GenericView struct {
	Name     string
	Fieldset []*Field
}

func NewGenericView(name string, fieldset []*Field) (me *GenericView) {
	me = &GenericView{
		Name:     name,
		Fieldset: fieldset,
	}

	return
}

func (me *GenericView) Publish(addr string, mode action.Action) error {
	var byt []byte
	var e error

	var form = "<form id=\"fEditor\" name=\"fEditor\">\n%s</form>\n"
	var hid = "\t<input type=\"hidden\" id=\"%s\" name=\"%s\" value=\"{{ %s|default('%s') }}\" />\n"
	var formgroup = "\t<div id=\"%sGroup\" class=\"form-group\" style=\"display: {{ %s|default('block') }};\">\n"
	formgroup += "\t\t<label for=\"%s\">%s</label>\n"
	formgroup += "\t\t<input type=\"%s\" class=\"form-control\" id=\"%s\" name=\"%s\" %s %s value=\"%s\" />\n"
	formgroup += "\t</div>\n"
	var formoptiongroup = "\t<div id=\"%sGroup\" class=\"form-group\" style=\"display: {{ %s|default('block') }};\">\n"
	formoptiongroup += "\t\t<label for=\"%s\">%s</label>\n"
	formoptiongroup += "\t\t<select class=\"form-control\" id=\"%s\" name=\"%s\" %s>\n"
	formoptiongroup += "\t\t\t<option value=\"\">--</option>\n"
	formoptiongroup += "\t\t\t{%% for option in enum_%d %%}\n"
	formoptiongroup += "\t\t\t<option value=\"{{ option.%s }}\" %s>{{ option.%s }}</option>\n"
	formoptiongroup += "\t\t\t{%% endfor %%}\n"
	formoptiongroup += "\t\t</select>\n"
	formoptiongroup += "\t</div>\n"

	var php = ""
	for i, field := range me.Fieldset {
		if field.AutoIncrement {
			php += fmt.Sprintf(hid, field.Name, field.Name, field.Name, field.Dv)
			continue
		}

		if field.Optionset == nil {
			php += fmt.Sprintf(formgroup, field.Name, field.Visibility(), field.Name, field.Star()+field.Description, field.Ft, field.Name, field.Name, field.Required(), field.MaxLength(), field.ValueByAction())
		} else {
			php += fmt.Sprintf(formoptiongroup, field.Name, field.Visibility(), field.Name, field.Star()+field.Description, field.Name, field.Name, field.Required(), i, field.Optionset.Key, field.SelectedByValue(), field.Optionset.Value)
		}
	}

	php = fmt.Sprintf(form, php)

	if _, e = os.Stat(fmt.Sprintf("%s\\views\\%s", addr, me.Name)); os.IsNotExist(e) {
		if e = os.Mkdir(fmt.Sprintf("%s\\views\\%s", addr, me.Name), os.ModePerm); e != nil {
			return e
		}
	}

	byt = []byte(php)
	if e = os.WriteFile(fmt.Sprintf("%s\\views\\%s\\%s.volt", addr, me.Name, mode.ToString()), byt, 0644); e != nil {
		return e
	}

	return e
}
