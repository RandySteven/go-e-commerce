package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
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

func BindJSON(req *http.Request, request any) error {
	return json.NewDecoder(req.Body).Decode(&request)
}

func ContentType(res http.ResponseWriter, contentType string) {
	res.Header().Set("Content-Type", contentType)
}
