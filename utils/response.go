package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseHandler(res http.ResponseWriter, code int, response any) {
	res.WriteHeader(code)
	json.NewEncoder(res).Encode(response)
}
