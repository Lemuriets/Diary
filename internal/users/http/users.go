package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lemuriets/diary/pkg/httpjson"
	"github.com/Lemuriets/diary/pkg/parseurl"
)

func (h *Handler) HelloMsg(w http.ResponseWriter, r *http.Request) {
	httpjson.InfoOKResponse(w, "This is an users service")
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseurl.GetIdUint(r)

	if err != nil {
		httpjson.NotFoundResponse(w)
		return
	}
	user := h.UseCase.GetById(uint(id))
	httpjson.OKResponse(w, user)

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseurl.GetIdUint(r)
	updateValues := map[string]interface{}{}

	if err != nil {
		httpjson.NotFoundResponse(w)
		return
	}
	if r.Form == nil {
		r.ParseMultipartForm(32 << 20)
	}
	for k, v := range r.Form {
		updateValues[k] = v[0]
	}
	h.UseCase.Update(id, updateValues)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseurl.GetIdUint(r)

	if err != nil {
		httpjson.NotFoundResponse(w)
		return
	}
	if err := h.UseCase.Delete(id); err != nil {
		httpjson.InternalServerErrorResponse(w)
		return
	}
	httpjson.RespondJSON(w, id, http.StatusOK)
}

func (h *Handler) MultipleDelete(w http.ResponseWriter, r *http.Request) {
	ids := r.FormValue("ids")

	if ids == "" {
		httpjson.BadRequestResponse(w)
		return
	}
	idsStrArray := strings.Split(ids, ",")
	idsArray := make([]uint, len(idsStrArray))

	for i, v := range idsStrArray {
		id, err := strconv.Atoi(v)

		if err != nil {
			httpjson.BadRequestResponse(w)
			return
		}
		idsArray[i] = uint(id)
	}
	fmt.Println(idsArray)
	if err := h.UseCase.MultipleDelete(idsArray); err != nil {
		httpjson.InternalServerErrorResponse(w)
		return
	}
	httpjson.InfoOKResponse(w, fmt.Sprintf("successful delete users with ids: %v", idsArray))
}

func (h *Handler) AddGrade(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r)
}
