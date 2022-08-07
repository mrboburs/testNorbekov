package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/mrboburs/Norbekov/model"
	"github.com/mrboburs/Norbekov/package/repository"
	"github.com/mrboburs/Norbekov/util/logrus"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

const (
	salt       = "hjqrhjqw124617aj564u564a654u65465aufhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353K554245987uaS?SFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims

	UserId int `json:"user_id"`
}

func (s *AdminService) GenerateToken(user_name, password string, logrus *logrus.Logger) (string, error) {
	user, err := s.repo.GetAdmin(user_name, password, logrus)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AdminService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AdminService) CreateAdmin(adminData model.Admin, logrus *logrus.Logger) (int, error) {
	return s.repo.CreateAdmin(adminData, logrus)
}

func (s *AdminService) GetAdmin(user_name, password string, logrus *logrus.Logger) (model.Admin, error) {
	return s.repo.GetAdmin(user_name, password, logrus)
}

func (s *AdminService) CheckAdmin(id int, logrus *logrus.Logger) (bool, error) {
	return s.repo.CheckAdmin(id, logrus)
}

func (s *AdminService) DeleteAdmin(id string, logrus *logrus.Logger) error {
	return s.repo.DeleteAdmin(id, logrus)
}
