package user_web

type AuthResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	Token string `json:"token"`
}