package usecase

import (
	"github.com/Lemuriets/diary/internal/auth"
)

type UseCase struct {
	Repository auth.Repository
}
