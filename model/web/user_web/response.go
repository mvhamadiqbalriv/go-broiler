package user_web

type Response struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}