package httpjson

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, payload interface{}, statusCode int) {
	response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(response))
}

func NotFoundResponse(w http.ResponseWriter) {
	result := map[string]string{
		"message": "404 not found",
	}
	RespondJSON(w, result, http.StatusNotFound)
}

func UnauthorizedResponse(w http.ResponseWriter) {
	result := map[string]string{
		"message": "401 unauthorized",
	}
	RespondJSON(w, result, http.StatusUnauthorized)
}

func InternalServerErrorResponse(w http.ResponseWriter) {
	result := map[string]string{
		"message": "500 internal server error",
	}
	RespondJSON(w, result, http.StatusInternalServerError)
}

func OKResponse(w http.ResponseWriter, payload interface{}) {
	RespondJSON(w, payload, http.StatusOK)
}

func InfoResponse(w http.ResponseWriter, message string) {
	result := map[string]string{
		"message": message,
	}
	OKResponse(w, result)
}
