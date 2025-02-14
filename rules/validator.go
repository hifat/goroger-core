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
	Validate(err error) (*map[string]string, error)
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

func (p *pgValidator) Validate(err error) (*map[string]string, error) {
	if _, ok := err.(validator.ValidationErrors); !ok {
		return nil, err
	}

	objErr := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		objErr[e.Field()] = e.Translate(p.trans)
	}

	return &objErr, nil
}
