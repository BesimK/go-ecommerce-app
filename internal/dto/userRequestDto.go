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

type SellerInput struct {
	FirstName         string `json:"first_name"        validate:"required"`
	LastName          string `json:"last_name"         validate:"required"`
	PhoneNumber       string `json:"phone_number"      validate:"required"`
	BankAccountNumber uint   `json:"bankAccountNumber" validate:"required"`
	SwiftCode         string `json:"swiftCode"         validate:"required"`
	PaymentType       string `json:"paymentType"       validate:"required"`
}
