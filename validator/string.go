package vl

import (
	"regexp"

	"github.com/Oudwins/zog"
)

type StringSchema struct {
	schema *zog.StringSchema
}

func String() *StringSchema {
	return &StringSchema{
		schema: zog.String(),
	}
}

func (v *StringSchema) Trim() *StringSchema {
	return &StringSchema{
		schema: v.schema.Trim(),
	}
}

func (v *StringSchema) Required() *StringSchema {
	return &StringSchema{
		schema: v.schema.Required(),
	}
}

func (v *StringSchema) Optional() *StringSchema {
	return &StringSchema{
		schema: v.schema.Optional(),
	}
}

func (v *StringSchema) Default(val string) *StringSchema {
	return &StringSchema{
		schema: v.schema.Default(val),
	}
}

func (v *StringSchema) Catch(val string) *StringSchema {
	return &StringSchema{
		schema: v.schema.Catch(val),
	}
}

func (v *StringSchema) OneOf(enum []string) *StringSchema {
	return &StringSchema{
		schema: v.schema.OneOf(enum),
	}
}

func (v *StringSchema) Min(n int) *StringSchema {
	return &StringSchema{
		schema: v.schema.Min(n),
	}
}

func (v *StringSchema) Max(n int) *StringSchema {
	return &StringSchema{
		schema: v.schema.Max(n),
	}
}

func (v *StringSchema) Len(n int) *StringSchema {
	return &StringSchema{
		schema: v.schema.Len(n),
	}
}

func (v *StringSchema) Email() *StringSchema {
	return &StringSchema{
		schema: v.schema.Email(),
	}
}

func (v *StringSchema) URL() *StringSchema {
	return &StringSchema{
		schema: v.schema.URL(),
	}
}

func (v *StringSchema) HasPrefix(s string) *StringSchema {
	return &StringSchema{
		schema: v.schema.HasPrefix(s),
	}
}

func (v *StringSchema) HasSuffix(s string) *StringSchema {
	return &StringSchema{
		schema: v.schema.HasSuffix(s),
	}
}

func (v *StringSchema) Contains(sub string) *StringSchema {
	return &StringSchema{
		schema: v.schema.Contains(sub),
	}
}

func (v *StringSchema) ContainsUpper() *StringSchema {
	return &StringSchema{
		schema: v.schema.ContainsUpper(),
	}
}

func (v *StringSchema) ContainsDigit() *StringSchema {
	return &StringSchema{
		schema: v.schema.ContainsDigit(),
	}
}

func (v *StringSchema) ContainsSpecial() *StringSchema {
	return &StringSchema{
		schema: v.schema.ContainsSpecial(),
	}
}

func (v *StringSchema) UUID() *StringSchema {
	return &StringSchema{
		schema: v.schema.UUID(),
	}
}

func (v *StringSchema) Match(regex *regexp.Regexp) *StringSchema {
	return &StringSchema{
		schema: v.schema.Match(regex),
	}
}
