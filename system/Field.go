package system

import (
	"defaultvalue"
	"fieldtype"
	"fmt"
)

type Field struct {
	Name          string
	Ft            fieldtype.FieldType
	Size          int
	Dv            defaultvalue.DefaultValue
	AllowNull     bool
	PrimaryKey    bool
	AutoIncrement bool
	Description   string
	Optionset     *Enum
}

func NewField(name string, ft fieldtype.FieldType, size int, dv defaultvalue.DefaultValue, allowNull bool, primaryKey bool, autoIncrement bool, description string, optionset *Enum) (me *Field) {
	me = &Field{
		Name:          name,
		Ft:            ft,
		Size:          size,
		Dv:            dv,
		AllowNull:     allowNull,
		PrimaryKey:    primaryKey,
		AutoIncrement: autoIncrement,
		Description:   description,
		Optionset:     optionset,
	}

	return
}

func (me *Field) Visibility() string {
	if me.PrimaryKey {
		return "identity"
	} else {
		return "'block'"
	}
}

func (me *Field) Star() string {
	if me.AllowNull {
		return ""
	} else {
		return "*"
	}
}

func (me *Field) Required() string {
	if me.AllowNull {
		return ""
	} else {
		return "required"
	}
}

func (me *Field) MaxLength() string {
	if me.Size != -1 {
		return fmt.Sprintf("maxlength=\"%d\"", me.Size)
	} else {
		return ""
	}
}

func (me *Field) ValueByAction() string {
	var value = "{{ %s|default('%s') }}"
	return fmt.Sprintf(value, me.Name, me.Dv)
}

func (me *Field) SelectedByValue() string {
	var selected = "{%% if option.%s === %s|default('%s') %%} selected {%% endif %%}"
	return fmt.Sprintf(selected, me.Optionset.Key, me.Name, me.Dv)
}
