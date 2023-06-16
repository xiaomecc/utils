package validator

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidate() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(r map[string]string, i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return errors.New(getErrorMsg(r, err))
	}
	return nil
}

func (cv *CustomValidator) Binder() {
	if err := cv.validator.RegisterValidation("mobile86", Mobile86); err != nil {
		log.Fatal("mobile rule reg error", err.Error())
	}
	if err := cv.validator.RegisterValidation("mobile880", Mobile880); err != nil {
		log.Fatal("mobile rule reg error", err.Error())
	}
}

func getErrorMsg(r map[string]string, err error) string {
	var ve validator.ValidationErrors
	var je *json.UnmarshalTypeError
	if errors.As(err, &ve) {
		for _, v := range err.(validator.ValidationErrors) {
			if message, exist := r[v.Field()+"."+v.Tag()]; exist {
				return message
			}
			return v.Error()
		}
	}
	if errors.As(err, &je) {
		return err.Error()
	}
	return "check error"
}

// Mobile mobile validate
var Mobile86 validator.Func = func(fl validator.FieldLevel) bool {
	m := fl.Field().String()
	pattern := `^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`
	ok, _ := regexp.MatchString(pattern, m)
	if !ok {
		return false
	}
	return true
}

var Mobile880 validator.Func = func(fl validator.FieldLevel) bool {
	m := fl.Field().String()
	pattern := `^(\+880)?(1){1}[56789]{1}(\d){8}`
	ok, _ := regexp.MatchString(pattern, m)
	if !ok {
		return false
	}
	return true
}
