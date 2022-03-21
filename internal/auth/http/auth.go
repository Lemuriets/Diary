package http

import (
	"net/http"

	"github.com/Lemuriets/diary/model"
	"github.com/Lemuriets/diary/pkg/crypto"
	"github.com/Lemuriets/diary/pkg/httpjson"
)

func (h *Handler) HelloMsg(w http.ResponseWriter, r *http.Request) {
	httpjson.RespondJSON(w, map[string]string{
		"message": "this is auth service",
	}, 200)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	token, err := h.UseCase.GenerateJWT(r.FormValue("login"), r.FormValue("password"))

	if err != nil {
		httpjson.RespondJSON(w, map[string]string{
			"message": "the user wasn't found",
		}, 401)
	} else {
		httpjson.RespondJSON(w, token, 200)
	}
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
		httpjson.RespondJSON(w, map[string]string{
			"message": err.Error(),
		}, 400)
	} else {
		httpjson.RespondJSON(w, user, 200)
	}
}
