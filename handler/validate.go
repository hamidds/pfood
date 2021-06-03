package handler

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
	"github.com/hamidds/pfood/model"
	"log"
)

func CustomerValidate(customer *model.Customer) error {
	validate, trans := SetUpValidator()
	filedErrors := validate.Struct(customer)
	return getError(filedErrors, trans)
}

func FoodValidate(food *model.Food) error {
	validate, trans := SetUpValidator()
	filedErrors := validate.Struct(food)
	return getError(filedErrors, trans)
}

func getError(filedErrors error, trans ut.Translator) error {
	if filedErrors != nil {
		var message string
		for _, e := range filedErrors.(validator.ValidationErrors) {
			fmt.Println(e.Translate(trans))
			message = e.Translate(trans)
		}
		return errors.New(message)
	}
	return nil
}



func SetUpValidator() (*validator.Validate, ut.Translator) {
	validate := validator.New()
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	_ = validate.RegisterTranslation("isdefault", trans, func(ut ut.Translator) error {
		return ut.Add("isdefault", "{0} should not be sent from user!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isdefault", fe.Field())

		return t
	})

	return validate, trans
}
