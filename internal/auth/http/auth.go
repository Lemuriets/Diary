package http

import (
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.UseCase.SignUp(w, r)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

}
