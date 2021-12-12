package service

import (
	"CinemaGCP/pkg/repository"
	"CinemaGCP/pkg/src"
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
)

const (
	salt          = "gsdgdsoh412412rnef"
	defaultBudget = 50
)

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (auth *AuthService) CreateUser(user src.User) (uuid.UUID, error) {
	user.Password = generatePasswordHash(user.Password)
	user.Budget = defaultBudget
	return auth.rep.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
