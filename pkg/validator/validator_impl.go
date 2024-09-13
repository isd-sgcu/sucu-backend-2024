package validator

import (
	"errors"

	"github.com/go-playground/locales/en"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type dtoValidator struct {
	v     *validator.Validate
	trans ut.Translator
}

func NewDtoValidator() (DTOValidator, error) {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil, errors.New("translator not found")
	}

	v := validator.New()
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil, err
	}

	return &dtoValidator{
		v:     v,
		trans: trans,
	}, nil
}

func (v *dtoValidator) Validate(dto interface{}) (errors []string) {
	err := v.v.Struct(dto)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Translate(v.trans))
		}
	}

	return errors
}
