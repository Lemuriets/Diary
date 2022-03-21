package parseurl

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIdUint64(r *http.Request) (uint64, error) {
	id := mux.Vars(r)["id"]

	return strconv.ParseUint(id, 10, 64)
}
