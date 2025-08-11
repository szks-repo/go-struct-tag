package tag

import "reflect"

type StructTag struct {
	tag reflect.StructTag
}

func NewFromField(field reflect.StructField) *StructTag {
	return &StructTag{
		tag: field.Tag,
	}
}
