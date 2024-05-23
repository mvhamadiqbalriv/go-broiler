package user_web

type UpdateRequest struct {
	Id	   int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}