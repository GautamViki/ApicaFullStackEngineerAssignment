package helper

import (
	"encoding/json"
	"net/http"
)

// Used for responding error messages in JSON content with w and code
func RespondWithError(w http.ResponseWriter, code int, msg string, responseCode string) {
	res := PrepareResponse(responseCode, msg)
	RespondwithJSON(w, code, res)
}


type ValidationErrors struct {
	ErrorMessages []string `json:"messages"`
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	ValidationErrors
}

// Prepare response
func PrepareResponse(code string, message string) Response {
	res := Response{}
	res.Code = code
	if message != "" {
		res.Message = message
	}
	res.ErrorMessages = []string{}
	return res
}

// Used for responding API with JSON content with w and code
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}