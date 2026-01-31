package validator

import (
	"errors"

	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type Validator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func NewValidator() *Validator {
	en := enLocale.New()
	uni := ut.New(en, en)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil
	}

	validate := validator.New()

	return &Validator{
		Validator:  validate,
		Translator: trans,
	}
}

func (v *Validator) Validate(data interface{}) error {
	err := v.Validator.Struct(data)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		for _, e := range validationErrors {
			log.Infof("[Validator] %s : %s", e.Field(), e.Translate(v.Translator))
			return errors.New(e.Translate(v.Translator))
		}
	}
	return nil
}
