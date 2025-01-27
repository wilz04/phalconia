package system

type Enum struct {
	Name  string
	Key   string
	Value string
}

func NewEnum(model *GenericModel, key *Field, value *Field) (me *Enum) {
	me = &Enum{
		Name:  model.Name + model.Suffix,
		Key:   key.Name,
		Value: value.Name,
	}

	return
}
