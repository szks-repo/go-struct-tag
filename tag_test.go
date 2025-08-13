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

func TestStructTag_Delimited(t *testing.T) {
	t.Parallel()

	type args struct {
		tagKey string
		delim  Delimiter
	}

	tets := []struct {
		anyStruct  any
		args       args
		want       *DelimitedValues
		itemChecks func(*testing.T, *DelimitedValues)
	}{
		{
			anyStruct: struct {
				Name string `json:"name,omitempty" form:"name"`
			}{},
			args: args{
				tagKey: "json",
				delim:  Delimiter{Delim: ","},
			},
			want: &DelimitedValues{
				delim: Delimiter{Delim: ","},
				values: []*DelimitedValue{
					{Key: "name"},
					{Key: "omitempty"},
				},
			},
		},
		{
			anyStruct: struct {
				Age int32 `json:"age,omitzero" gorm:"index:,class:FULLTEXT,comment:hello world,where:age > 10"`
			}{},
			args: args{
				tagKey: "gorm",
				delim:  Delimiter{Delim: ",", KeyValueSep: ":"},
			},
			want: &DelimitedValues{
				delim: Delimiter{Delim: ",", KeyValueSep: ":"},
				values: []*DelimitedValue{
					{Key: "index"},
					{Key: "class", Value: "FULLTEXT"},
					{Key: "comment", Value: "hello world"},
					{Key: "where", Value: "age > 10"},
				},
			},
			itemChecks: func(t *testing.T, delimiteds *DelimitedValues) {
				assert.True(t, delimiteds.HasKey("index"))
				assert.True(t, delimiteds.HasKey("class"))
				assert.True(t, delimiteds.HasKey("comment"))
				assert.True(t, delimiteds.HasKey("where"))
				assert.False(t, delimiteds.HasKey("none"))
			},
		},
	}

	for _, tt := range tets {
		t.Run("", func(t *testing.T) {
			tag := NewTagFromField(reflect.TypeOf(tt.anyStruct).Field(0))
			item, _ := tag.Get(tt.args.tagKey)
			got := item.Delimited(tt.args.delim)
			assert.Equal(t, tt.want, got)
			if tt.itemChecks != nil {
				tt.itemChecks(t, got)
			}
		})
	}
}
