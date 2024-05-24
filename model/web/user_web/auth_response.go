package user_web

type LoggedInUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}

type AuthResponse struct {
	LoggedInUser LoggedInUser `json:"logged_in_user"`
	Token        string       `json:"token"`
}