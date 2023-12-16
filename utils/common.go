package utils

import (
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
