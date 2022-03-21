package http

import (
	"fmt"
	"net/http"

	"github.com/Lemuriets/diary/pkg/httpjson"
	"github.com/Lemuriets/diary/pkg/parseurl"
)

func (h *Handler) HelloMsg(w http.ResponseWriter, r *http.Request) {
	httpjson.OKResponse(w, "This is an users service")
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseurl.GetIdUint64(r)

	if err != nil {
		httpjson.NotFoundResponse(w)
		return
	}
	user := h.UseCase.GetById(id)
	httpjson.OKResponse(w, user)

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	_, err := parseurl.GetIdUint64(r)

	if err != nil {
		httpjson.NotFoundResponse(w)
		return
	}
	// I'll do it later
	fmt.Println(r.PostForm)
	// h.UseCase.Update(id, r.PostForm)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseurl.GetIdUint64(r)

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
