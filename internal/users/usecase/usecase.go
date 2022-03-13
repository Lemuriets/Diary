package usecase

import (
	"github.com/Lemuriets/diary/internal/users"
)

type UseCase struct {
	Repository users.Repository
}
