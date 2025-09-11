package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = TranslateTag(v)
		}
	}
	return res
}

func TranslateTag(fd validator.FieldError) string {
	switch fd.ActualTag() {
	case "required" :
		return fmt.Sprintf("fields %s is required", fd.StructField())
	case "email" :
		return fmt.Sprintf("fields %s is must be a valid email", fd.Field())
	case "min" :
		return fmt.Sprintf("fielfs %s must be at least %s characters", fd.Field(), fd.Param())
	case "max" :
		return  fmt.Sprintf("fields %s must be at most %s characters", fd.Field(), fd.Param())
	}
	return "Validation error"
}