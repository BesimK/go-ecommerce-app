package service

import "github.com/BesimK/go-ecommerce-app/internal/domain"

type UserService struct{}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (s UserService) Signup(input any) (string, error) {
	return "", nil
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

func (s UserService) FindCart(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) CreateCart(e domain.User) (int, error) {
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
