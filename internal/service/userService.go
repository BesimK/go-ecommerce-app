package service

import (
	"fmt"

	"github.com/BesimK/go-ecommerce-app/internal/domain"
	"github.com/BesimK/go-ecommerce-app/internal/dto"
)

type UserService struct{}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	return "this-is-my-token-as-of-now", nil
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	fmt.Println(email)
	return nil, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
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

func (s UserService) FindCart([]interface{}) (int, error) {
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
