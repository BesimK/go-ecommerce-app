package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/BesimK/go-ecommerce-app/config"
	"github.com/BesimK/go-ecommerce-app/internal/domain"
	"github.com/BesimK/go-ecommerce-app/internal/dto"
	"github.com/BesimK/go-ecommerce-app/internal/helper"
	"github.com/BesimK/go-ecommerce-app/internal/repository"
	"github.com/BesimK/go-ecommerce-app/pkg/notification"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	hPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", nil
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})
	if err != nil {
		log.Println("This user exists")
		return "", errors.New("this user already exist")
	}
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)
	if err != nil {
		log.Printf("error finding user by email %v", err)
		return nil, errors.New("had problem")
	}

	return &user, nil
}

func (s UserService) isVerifiedUser(id uint) bool {
	currentUser, err := s.Repo.FindUserByID(id)

	return currentUser.Verified && err == nil
}

func (s UserService) GetVerificationCode(e domain.User) error {
	if s.isVerifiedUser(e.ID) {
		return errors.New("already verified")
	}

	code, err := s.Auth.GenerateCode()
	if err != nil {
		return err
	}

	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	user, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil {
		return errors.New("unable to update verication code")
	}

	notificationClient := notification.NewNotificationClient(s.Config)

	msg := fmt.Sprintf("Your verification code is here baby %v", code)

	if err = notificationClient.SendSMS(user.Phone, msg); err != nil {
		return err
	}

	return nil
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) VerifyCode(id uint, code int) error {
	if s.isVerifiedUser(id) {
		return errors.New("user is already verified")
	}

	user, err := s.Repo.FindUserByID(id)
	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if time.Now().After(user.Expiry) {
		return errors.New("verification code expired")
	}

	if _, err := s.Repo.UpdateUser(id, domain.User{
		Verified: true,
	}); err != nil {
		return errors.New("unable to update user")
	}

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(input any) (string, error) {
	return "", nil
}

func (s UserService) BecomeSeller(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) FindCart([]any) (int, error) {
	return 0, nil
}

func (s UserService) CreateCart(input any, u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) CreateOrder(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrderByID(e domain.User) (int, error) {
	return 0, nil
}
