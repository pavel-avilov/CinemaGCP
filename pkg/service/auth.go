package service

import (
	"CinemaGCP/pkg/logger"
	"CinemaGCP/pkg/repository"
	"CinemaGCP/pkg/src"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const (
	salt          = "gsdgdsoh412412rnef"
	defaultBudget = 50
	signingKey    = "gfgdf234212fFSD#123fsd"
	tokenTime     = 5 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id"`
}

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

func (auth *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := auth.rep.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	logger.Info(user.Id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
