package usecase

import (
	"log"
	"net/http"

	"github.com/Lemuriets/diary/internal/auth"
	"github.com/Lemuriets/diary/pkg/httpjson"
)

type UseCase struct {
	Repository auth.Repository
}

func NewUseCase(repo auth.Repository) *UseCase {
	return &UseCase{
		Repository: repo,
	}
}

func (uc *UseCase) SignIn(w http.ResponseWriter, r *http.Request) {
	countUsers := uc.Repository.GetCount(r.FormValue("login"), r.FormValue("password"))

	if countUsers == 0 {

	}

}

func (uc *UseCase) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := uc.Repository.Get(r.FormValue("login"), r.FormValue("password"))

	if err != nil {
		log.Fatal(err)
	}

	httpjson.RespondJSON(w, 200, user)
}
