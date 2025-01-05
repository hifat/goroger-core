// package validator

// import "github.com/go-playground/validator"

// type User struct {
// 	Username string `json:"username" validate:"required"`
// 	Password string `json:"password" validate:"required"`
// }

// func (u *User) Validate() error {
// 	var validate *validator.Validate
// 	err := validate.Struct(u)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

package validator

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type IValidator interface {
	ValidateStruct(interface{}) error
}

type goValidator struct {
	validate *validator.Validate
}

func NewGoValidator() IValidator {
	v := validator.New()
	// Register custom tag name
	v.SetTagName("rule")
	// Register custom validation rule
	err := v.RegisterValidation("important", validator.Func(func(fl validator.FieldLevel) bool {
		return fl.Field().Len() > 0
	}))
	if err != nil {
		log.Fatal(err)
	}

	return &goValidator{
		validate: v,
	}
}

func (v *goValidator) ValidateStruct(s interface{}) error {
	return v.validate.Struct(s)
}

type User struct {
	Username string `json:"username" rule:"important"`
	Password string `json:"password" rule:"important"`
}

func (u *User) Validate(v IValidator) error {
	return v.ValidateStruct(u)
}
