package tag

import (
	"reflect"
	"strings"
)

// StructTag represents a collection of tags for struct fields.
//
//	struct User {
//	    Name string `json:"name,omitempty" form:"name"`
//	}
//
// It allows for easy access to tags by namespace, such as "json" or "form".
type StructTag struct {
	raw   reflect.StructTag
	items []Item
}

func (t *StructTag) Get(key string) (*Item, bool) {
	for _, item := range t.items {
		if item.Key == key {
			return &item, true
		}
	}
	return nil, false
}

func (t *StructTag) Delete(key string) {
	for i, item := range t.items {
		if item.Key == key {
			t.items = append(t.items[:i], t.items[i+1:]...)
			return
		}
	}
}

func NewTagFromField(field reflect.StructField) *StructTag {
	st := &StructTag{
		raw:   field.Tag,
		items: make([]Item, 0),
	}
	if field.Tag != "" {
		st.items = parse(string(field.Tag))
	}
	return st
}

func parse(s string) []Item {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}

	var items []Item
	for s != "" {
		var i int
		if i < len(s) && s[i] == ' ' {
			i++
		}
		s = s[i:]
		if s == "" {
			break
		}

		i = 0
		for i < len(s) && s[i] != ':' && s[i] != ' ' && s[i] != '"' {
			i++
		}
		if i == 0 {
			break
		}

		key := s[:i]
		s = s[i:]

		if len(s) == 0 || s[0] != ':' {
			items = append(items, Item{Key: key})
			continue
		}
		// skip ':'
		s = s[1:]

		if len(s) > 0 && s[0] == '"' {
			// skip opening quote
			s = s[1:]
			value, remaining := parseQuotedString(s)
			items = append(items, Item{Key: key, Value: value})
			s = remaining
		} else {
			i = 0
			for i < len(s) && s[i] != ' ' {
				i++
			}
			value := s[:i]
			items = append(items, Item{Key: key, Value: value})
			s = s[i:]
		}
	}

	return items
}

func parseQuotedString(s string) (string, string) {
	var buf strings.Builder
	var escaped bool
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if escaped {
			switch ch {
			case '"', '\\', 'n', 'r', 't':
				buf.WriteByte(s[i])
			default:
				buf.WriteByte('\\')
				buf.WriteByte(s[i])
			}
			escaped = false
		} else if s[i] == '\\' {
			escaped = true
		} else if s[i] == '"' {
			return buf.String(), s[i+1:] // return the string and the remaining part
		} else {
			buf.WriteByte(s[i])
		}
	}

	return buf.String(), ""
}

func escape(s string) string {
	var buf strings.Builder
	for _, c := range s {
		switch c {
		case '"':
			buf.WriteString("\\\"")
		case '\\':
			buf.WriteString("\\\\")
		case '\n':
			buf.WriteString("\\n")
		case '\r':
			buf.WriteString("\\r")
		case '\t':
			buf.WriteString("\\t")
		default:
			buf.WriteRune(c)
		}
	}
	return s
}

type Item struct {
	Key   string // e.g. "json", "form"
	Value string
}

func (i *Item) Delimited(delim Delimiter) *DelimitedValues {
	if delim.Delim == "" {
		return &DelimitedValues{
			delim:  delim,
			values: []*DelimitedValue{{Key: i.Value}},
		}
	}

	// TODO: consider escaping the delimiter in the value
	var delimiteds []*DelimitedValue
	for _, delimited := range strings.Split(i.Value, delim.Delim) {
		if delim.KeyValueSep != "" && strings.Contains(delimited, delim.KeyValueSep) {
			parts := strings.SplitN(delimited, delim.KeyValueSep, 2)
			if len(parts) == 2 {
				delimiteds = append(delimiteds, &DelimitedValue{
					Key:   parts[0],
					Value: parts[1],
				})
			}
		} else {
			delimiteds = append(delimiteds, &DelimitedValue{
				Key: delimited,
			})
		}
	}

	return &DelimitedValues{
		values: delimiteds,
		delim:  delim,
	}
}

type Delimiter struct {
	Delim       string
	KeyValueSep string
}

type DelimitedValue struct {
	Key   string
	Value string
}

type DelimitedValues struct {
	delim  Delimiter
	values []*DelimitedValue
}

func (d DelimitedValues) HasKey(key string) bool {
	for _, item := range d.values {
		if item.Key == key {
			return true
		}
	}
	return false
}

func (d DelimitedValues) Get(key string) (*DelimitedValue, bool) {
	for _, item := range d.values {
		if item.Key == key {
			return item, true
		}
	}
	return nil, false
}
