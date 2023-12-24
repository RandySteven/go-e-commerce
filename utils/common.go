package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/RandySteven/go-e-commerce.git/apperror"
	"github.com/go-playground/validator"
)

func GetModelName(m any) string {
	t := reflect.TypeOf(m)
	model := strings.ToLower(t.Name())
	return model
}

func GetFields(m any) map[string]any {
	t := reflect.TypeOf(m)
	result := make(map[string]any)
	v := reflect.ValueOf(m)
	for i := 0; i < t.NumField(); i++ {
		result[t.Field(i).Name] = v.Field(i)
	}
	return result
}

func DateValidation(birthdate time.Time) error {
	by, _, _ := birthdate.Date()
	today := time.Now()
	ty, _, _ := today.Date()
	age := ty - by
	if age < 18 {
		return &apperror.ErrAgeValidation{}
	}
	return nil
}

func BindJSON(req *http.Request, request any) error {
	return json.NewDecoder(req.Body).Decode(&request)
}

func ContentType(res http.ResponseWriter, contentType string) {
	res.Header().Set("Content-Type", contentType)
}

func Validate(req any) error {
	validate := validator.New()
	type requestLength struct {
		Min uint
		Max uint
	}

	requestLengthMap := map[string]requestLength{
		"Password": {
			Min: 8,
			Max: 16,
		},
	}

	err := validate.Struct(req)
	if err != nil {
		for _, currErr := range err.(validator.ValidationErrors) {
			switch currErr.ActualTag() {
			case "email":
				err = fmt.Errorf("%s field is not in email format", currErr.Field())
			case "min":
				err = fmt.Errorf("%s field is less than %d", currErr.Field(), requestLengthMap[currErr.Field()].Min)
			case "max":
				err = fmt.Errorf("%s field is more than %d", currErr.Field(), requestLengthMap[currErr.Field()].Max)
			case "numeric":
				err = fmt.Errorf("%s field must in numeric value ", currErr.Field())
			default:
				err = fmt.Errorf("%s field is %s", currErr.Field(), currErr.ActualTag())
			}
		}
		return err
	}
	return nil
}
