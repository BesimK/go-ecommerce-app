package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/BesimK/go-ecommerce-app/internal/domain"
)

type UserRepository interface {
	CreateUser(usr domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)

	// more functions will be shown!!
	CreateBankAccount(e domain.BankAccount) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {
	if err := r.db.Create(&usr).Error; err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return usr, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		log.Printf("findUser error  %v", err)

		return domain.User{}, errors.New("failed to find user")
	}

	return user, nil
}

func (r userRepository) FindUserByID(id uint) (domain.User, error) {
	var user domain.User

	if err := r.db.First(&user, id).Error; err != nil {
		log.Printf("findUserByIByIdd error  %v", err)

		return domain.User{}, errors.New("failed to find user by ID")
	}

	return user, nil
}

func (r userRepository) UpdateUser(
	id uint,
	u domain.User,
) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).
		Clauses(clause.Returning{}).
		Where("id=?", id).
		Updates(u).
		Error
	if err != nil {
		log.Printf("error on update %v", err)
		return domain.User{}, errors.New("failed update user")
	}

	return user, nil
}

func (r userRepository) CreateBankAccount(e domain.BankAccount) error {
	return r.db.Create(&e).Error
}
