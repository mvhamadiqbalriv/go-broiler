package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// CustomValidator holds the validator instance
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator creates a new instance of CustomValidator
func NewCustomValidator() *CustomValidator {
	cv := &CustomValidator{
		validator: validator.New(),
	}

	// Register custom validation functions
	cv.RegisterCustomValidation("confirmNewPassword", confirmNewPassword)

	return cv
}

// ValidateStruct validates the struct using the validator
func (cv *CustomValidator) ValidateStruct(s interface{}) error {
	if err := cv.validator.Struct(s); err != nil {
		return err
	}
	return nil
}

// RegisterCustomValidation registers a custom validation function
func (cv *CustomValidator) RegisterCustomValidation(tag string, fn validator.Func) {
	cv.validator.RegisterValidation(tag, fn)
}

// Custom validation function
func customValidation(fl validator.FieldLevel) bool {
	// Custom validation logic
	value := fl.Field().String()
	return value == "custom" // Example validation logic: value must be "custom"
}

// Confirm New Password
func confirmNewPassword(fl validator.FieldLevel) bool {
	newPassword := fl.Parent().FieldByName("NewPassword").String()
	confirmPassword := fl.Parent().FieldByName("ConfirmPassword").String()

	fmt.Println(newPassword, confirmPassword)
	return newPassword == confirmPassword
}