package user_web

// ChangeProfilePictureRequest represent the request that needed to change password
type CreateProfilePictureRequest struct {
	Id int `json:"id" validate:"required"`
	ProfilePicture string `json:"profile_picture" validate:"required,base64"`
}