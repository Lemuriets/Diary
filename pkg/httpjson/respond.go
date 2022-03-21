package httpjson

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, payload interface{}, statusCode int) {
	response, err := json.Marshal(payload)

	w.WriteHeader(statusCode)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
