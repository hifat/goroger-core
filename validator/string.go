package vl

import (
	"regexp"

	"github.com/Oudwins/zog"
	"github.com/Oudwins/zog/internals"
)

type stringSchema struct {
	schema *zog.StringSchema
}

func String() *stringSchema {
	return &stringSchema{
		schema: zog.String(),
	}
}

func (v *stringSchema) getSchema() zog.ZogSchema {
	return v.schema
}

func (v *stringSchema) Trim() *stringSchema {
	return &stringSchema{
		schema: v.schema.Trim(),
	}
}

func (v *stringSchema) setOptions(options ...VlTestOption) (zopts []func(test *internals.Test)) {
	for _, opt := range options {
		var vltest VlTest
		opt(&vltest)
		if vltest.ErrMessage != "" {
			zopts = append(zopts, zog.Message(vltest.ErrMessage))
		}
	}

	return zopts
}

func (v *stringSchema) Required(options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.Required(zopts...),
	}
}

func (v *stringSchema) Optional() *stringSchema {
	return &stringSchema{
		schema: v.schema.Optional(),
	}
}

func (v *stringSchema) Default(val string) *stringSchema {
	return &stringSchema{
		schema: v.schema.Default(val),
	}
}

func (v *stringSchema) Catch(val string) *stringSchema {
	return &stringSchema{
		schema: v.schema.Catch(val),
	}
}

func (v *stringSchema) OneOf(enum []string, options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.OneOf(enum, zopts...),
	}
}
func (v *stringSchema) Min(n int, options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.Min(n, zopts...),
	}
}

func (v *stringSchema) Max(n int, options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.Max(n, zopts...),
	}
}

func (v *stringSchema) Len(n int, options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.Len(n, zopts...),
	}
}

func (v *stringSchema) Email(options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.Email(zopts...),
	}
}

func (v *stringSchema) URL(options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.URL(zopts...),
	}
}

func (v *stringSchema) HasPrefix(s string, options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.HasPrefix(s, zopts...),
	}
}

func (v *stringSchema) HasSuffix(s string, options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.HasSuffix(s, zopts...),
	}
}

func (v *stringSchema) Contains(sub string, options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.Contains(sub, zopts...),
	}
}

func (v *stringSchema) ContainsUpper(options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.ContainsUpper(zopts...),
	}
}

func (v *stringSchema) ContainsDigit(options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.ContainsDigit(zopts...),
	}
}

func (v *stringSchema) ContainsSpecial(options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.ContainsSpecial(zopts...),
	}
}

func (v *stringSchema) UUID(options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.UUID(zopts...),
	}
}

func (v *stringSchema) Match(regex *regexp.Regexp, options ...VlTestOption) *stringSchema {
	zopts := v.setOptions(options...)
	return &stringSchema{
		schema: v.schema.Match(regex, zopts...),
	}
}
