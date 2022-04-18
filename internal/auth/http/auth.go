package http

import (
	"net/http"

	"github.com/Lemuriets/diary/internal/auth"
	"github.com/Lemuriets/diary/model"
	"github.com/Lemuriets/diary/pkg/crypto"
	"github.com/Lemuriets/diary/pkg/httpjson"
	"gorm.io/gorm"
)

func (h *Handler) HelloMsg(w http.ResponseWriter, r *http.Request) {
	httpjson.InfoResponse(w, "this is an auth service")
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	accessToken, errAc := h.UseCase.GenerateAccessToken(r.FormValue("login"), r.FormValue("password"))
	refreshToken, errRef := h.UseCase.GenerateRefreshToken(r.FormValue("login"))

	if errRef != nil || errAc == gorm.ErrRecordNotFound {
		httpjson.NotFoundResponse(w)
		return
	} else if errAc == auth.WrongPassword {
		httpjson.UnauthorizedResponse(w)
		return
	}
	httpjson.OKResponse(w, map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
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

func (h *Handler) UpdateTokens(w http.ResponseWriter, r *http.Request) {

}
