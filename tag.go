package tag

import "reflect"

// StructTags represents a collection of tags for struct fields.
//
//	struct User {
//	    Name string `json:"name,omitempty" form:"name"`
//	}
//
// It allows for easy access to tags by namespace, such as "json" or "form".
type StructTags struct {
	tag        reflect.StructTag
	namespaces map[string]string
}

func NewTagsFromField(field reflect.StructField) *StructTags {
	return &StructTags{
		tag:        field.Tag,
		namespaces: make(map[string]string),
	}
}
