package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/BesimK/go-ecommerce-app/internal/domain"
	"github.com/BesimK/go-ecommerce-app/internal/dto"
	"github.com/BesimK/go-ecommerce-app/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	fmt.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})
	if err != nil {
		fmt.Println("Existing User")
	}

	// generate token
	log.Println(user)

	userInfo := fmt.Sprintf(
		"%v, %v, %v, %v",
		user.ID,
		user.Email,
		user.Email,
		user.UserType,
	)

	return userInfo, nil
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)
	if err != nil {
		log.Printf("error finding user by email %v", err)
		return nil, errors.New("had problem")
	}

	return &user, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	// compate password and generate token
	return user.Email, nil
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
