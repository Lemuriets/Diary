package http

import (
	"net/http"

	"github.com/Lemuriets/diary/pkg/httpjson"
)

func (h *Handler) HelloMsg(w http.ResponseWriter, r *http.Request) {
	httpjson.RespondJSON(w, map[string]string{
		"message": "this is users service",
	}, 200)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	user := h.UseCase.GetById(1)

	httpjson.RespondJSON(w, user, 200)
}
