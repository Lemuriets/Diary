package http

import (
	"net/http"

	"github.com/Lemuriets/diary/model"
	"github.com/Lemuriets/diary/pkg/crypto"
	"github.com/Lemuriets/diary/pkg/httpjson"
)

func (h *Handler) HelloMsg(w http.ResponseWriter, r *http.Request) {
	httpjson.InfoResponse(w, "this is an auth service")
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	token, err := h.UseCase.GenerateJWT(r.FormValue("login"), r.FormValue("password"))

	if err != nil {
		httpjson.NotFoundResponse(w)
		return
	}
	httpjson.OKResponse(w, map[string]string{
		"token": token,
	})

}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := model.User{
		Login:        r.FormValue("login"),
		PasswordHash: crypto.HashPassword((r.FormValue("password"))),
		Name:         r.FormValue("name"),
		Lastname:     r.FormValue("lastname"),
		Patronymic:   r.FormValue("patronymic"),
	}

	if err := h.UseCase.Create(user); err != nil {
		httpjson.InternalServerErrorResponse(w)
		return
	}
	httpjson.RespondJSON(w, user, http.StatusOK)

}
