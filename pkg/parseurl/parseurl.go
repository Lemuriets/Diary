package parseurl

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIdUint(r *http.Request) (uint, error) {
	id := mux.Vars(r)["id"]
	result, err := strconv.ParseUint(id, 10, 64)

	return uint(result), err
}
