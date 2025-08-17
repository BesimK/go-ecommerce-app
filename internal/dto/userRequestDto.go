// Package dto
package dto

type UserLogin struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserSignup struct {
	UserLogin
	Phone string `json:"phone" validate:"required"`
}

type VerificationCodeInput struct {
	Code int `json:"code" validate:"required"`
}
