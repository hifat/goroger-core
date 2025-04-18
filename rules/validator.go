package rules

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator interface {
	Validate(v any) (*map[string]string, error)
}

type pgValidator struct {
	v     *validator.Validate
	trans ut.Translator
}

func New(v *validator.Validate) Validator {
	return &pgValidator{v: v}
}

func Register() (*validator.Validate, error) {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}

		return name
	})

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	return v, en_translations.RegisterDefaultTranslations(v, trans)
}

func (p *pgValidator) Struct(v any) error {
	return p.v.Struct(v)
}

func (p *pgValidator) Validate(v any) (*map[string]string, error) {
	err := p.v.Struct(v)
	validateErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, err
	}

	objErr := make(map[string]string)
	for _, e := range validateErrors {
		objErr[e.Field()] = e.Translate(p.trans)
	}

	return &objErr, nil
}
