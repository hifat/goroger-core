package vl

import (
	"fmt"
	"reflect"

	"github.com/Oudwins/zog"
)

type VlSchema interface {
	getSchema() zog.ZogSchema
}

type Schema map[string]VlSchema

type VlError struct {
	Attributes map[string][]string
	Message    string
}

func (e VlError) Error() string {
	return e.Message
}

func Parse[T any](data T, rules Schema) error {
	zschema := zog.Schema{}

	for k, v := range rules {
		zschema[k] = v.getSchema()
	}

	errs := zog.Struct(zschema).Parse(data, zschema)
	if errs != nil {
		mapErr := zog.Errors.SanitizeMap(errs)

		v := reflect.ValueOf(errs)
		keys := v.MapKeys()

		return VlError{
			Message:    fmt.Sprintf("%s %s", keys[0], errs["$first"][0].Message()),
			Attributes: mapErr,
		}
	}

	return nil
}
