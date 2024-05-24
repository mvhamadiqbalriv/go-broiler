package helper

import (
	"mvhamadiqbalriv/belajar-golang-restful-api/model/domain"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/web/user_web"
)

func ToUserResponse(user domain.User) user_web.Response {
	profilePicture := BaseUrl(user.ProfilePicture)
	return user_web.Response{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		ProfilePicture: profilePicture,
	}
}

func ToUserResponses(users []domain.User) []user_web.Response {
    var userResponses []user_web.Response
    for _, user := range users {
        userResponse := ToUserResponse(user)
		userResponses = append(userResponses, userResponse)
    }
    return userResponses
}

func ToUserResponseAuth(user domain.User, token string) user_web.AuthResponse {
	return user_web.AuthResponse{
		LoggedInUser: user_web.LoggedInUser{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			ProfilePicture: user.ProfilePicture,
		},
		Token: token,
	}
}