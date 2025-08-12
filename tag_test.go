package tag

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFromField(t *testing.T) {
	t.Parallel()

	tests := []struct {
		field func() reflect.StructField
		want  *StructTag
	}{
		{
			field: func() reflect.StructField {
				return reflect.TypeOf(struct {
					Name string `json:"name,omitempty" form:"name"`
				}{}).Field(0)
			},
			want: &StructTag{
				raw: reflect.StructTag(`json:"name,omitempty" form:"name"`),
				items: []Item{
					{Key: "json", Value: "name,omitempty"},
					{Key: "form", Value: "name"},
				},
			},
		},
		{
			field: func() reflect.StructField {
				return reflect.TypeOf(struct {
					Age int32 `gorm:"index:,class:FULLTEXT,comment:hello \\, world,where:age > 10"`
				}{}).Field(0)
			},
			want: &StructTag{
				raw: reflect.StructTag(`gorm:"index:,class:FULLTEXT,comment:hello \\, world,where:age > 10"`),
				items: []Item{
					{Key: "gorm", Value: "index:,class:FULLTEXT,comment:hello \\, world,where:age > 10"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := NewTagFromField(tt.field())
			assert.Equal(t, tt.want, got)
		})
	}
}
