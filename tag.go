package tag

import "reflect"

type StructTags struct {
	tag        reflect.StructTag
	namespaces map[string]string
}

func NewFromField(field reflect.StructField) *StructTags {
	return &StructTags{
		tag:        field.Tag,
		namespaces: make(map[string]string),
	}
}
