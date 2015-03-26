package util

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jadengore/Ricetta/api/types"
)

//
// Library of responses for the Ricetta API.
//

func (u Util) SimpleJsonResponse(w rest.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.WriteJson(types.Json{
		"response": message,
	})
}

func (u Util) SimpleJsonReason(w rest.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.WriteJson(types.Json{
		"reason": message,
	})
}

//
// Validator Responses
//

func (u Util) SimpleJsonValidationReason(w rest.ResponseWriter, code int, err []error) {
	errorMessage := decodeValidatorErrors(err)
	w.WriteHeader(code)
	w.WriteJson(types.Json{
		"reason": errorMessage,
	})
}

func decodeValidatorErrors(err []error) []string {
	errorMessage := make([]string, len(err))
	for i := range err {
		errorMessage[i] = err[i].Error()
	}
	return errorMessage
}

func (u Util) FailedToAuthenticate(w rest.ResponseWriter) {
	w.WriteHeader(401)
	w.WriteJson(types.Json{
		"response": "Failed to authenticate request",
		"reason":   "Missing, illegal or expired token",
	})
}
