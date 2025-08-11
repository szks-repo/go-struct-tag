package tag

import (
	"reflect"
)

// StructTag represents a collection of tags for struct fields.
//
//	struct User {
//	    Name string `json:"name,omitempty" form:"name"`
//	}
//
// It allows for easy access to tags by namespace, such as "json" or "form".
type StructTag struct {
	tag        reflect.StructTag
	namespaces map[string]Item
}

func NewTagsFromField(field reflect.StructField) *StructTag {
	st := &StructTag{
		tag:        field.Tag,
		namespaces: make(map[string]Item),
	}
	if field.Tag != "" {

	}
	return st
}

type Item struct {
	namespace string
}

func parseTag(s string) (string, Item) {
	return "", Item{}
}
