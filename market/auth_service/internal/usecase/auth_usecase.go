package usecase

import (
	"errors"
	"marketapi/auth/internal/crypto"
	repo "marketapi/auth/internal/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUseCase struct {
	r         repo.AuthRepo
	tokenTTl  time.Duration
	jwtSecret []byte
}

func New(repositorie repo.AuthRepo, jwt string) *AuthUseCase {
	return &AuthUseCase{
		r:         repositorie,
		tokenTTl:  time.Minute * 45,
		jwtSecret: []byte(jwt),
	}
}

func (us *AuthUseCase) TryLogin(login string, pwd string) (string, error) {
	user, err := us.r.TryLogin(login)
	if err != nil {
		return "", errors.New("user not found")
	}

	res := crypto.CheckPasswordHash(pwd, user.PasswordHash)
	if !res {
		return "", errors.New("wrong password")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(us.tokenTTl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(us.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
