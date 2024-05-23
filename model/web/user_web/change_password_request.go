package user_web

// ChangePasswordRequest represent the request that needed to change password
type ChangePasswordRequest struct {
	Id int `json:"id" validate:"required"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,confirmNewPassword"`
}