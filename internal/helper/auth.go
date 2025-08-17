package helper

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/BesimK/go-ecommerce-app/internal/domain"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(p string) (string, error) {
	if len(p) < 6 {
		return "", errors.New("password length must be at least 6")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return "", errors.New("error generating password")
	}

	return string(hashP), nil
}

func (a Auth) GenerateToken(
	id uint,
	email string,
	role string,
) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("required inputs to generate token are missing")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString(
		[]byte("a-string-secret-at-least-256-bits-long"),
	)
	if err != nil {
		return "", errors.New("error signing string")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(
	pP string,
	hP string,
) error {
	if len(pP) < 6 {
		return errors.New("password length must be more than 6")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP)); err != nil {
		log.Println("password does not match")
		return errors.New("password does not match")
	}

	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {
	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, errors.New("invalid token format")
	}

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token type")
	}

	tokenStr := tokenArr[1]
	token, err := jwt.Parse(
		tokenStr,
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"unknown signing method %v",
					token.Header,
				)
			}
			return []byte(a.Secret), nil
		},
	)
	if err != nil {
		return domain.User{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		return domain.User{
			ID:       uint(claims["user_id"].(float64)),
			Email:    claims["email"].(string),
			UserType: claims["role"].(string),
		}, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]
	if len(authHeader) == 0 {
		return errors.New("not provided token")
	}

	user, err := a.VerifyToken(authHeader[0])
	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "autorization failed",
			"reason":  err.Error(),
		})
	}
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}

func (a Auth) GenerateCode() (int, error) {
	return RandomNumbers(6)
}
