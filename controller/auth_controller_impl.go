package controller

import (
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/web"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/web/user_web"
	"mvhamadiqbalriv/belajar-golang-restful-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	UserService service.UserService
	TokenBlacklistService service.TokenBlacklistService
}

func NewAuthController(userService service.UserService, TokenBlacklistService service.TokenBlacklistService) AuthController {
	return &AuthControllerImpl{
		UserService: userService,
		TokenBlacklistService: TokenBlacklistService,
	}
}

func (controller *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userLoginRequest := user_web.LoginRequest{}
	helper.ReadFromRequestBody(r, &userLoginRequest)

	authResponse := controller.UserService.Login(r.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *AuthControllerImpl) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	tokenString := helper.ExtractTokenFromHeader(r)
	err := controller.TokenBlacklistService.AddTokenToBlacklist(tokenString)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   tokenString,
	}

	helper.WriteToResponseBody(w, webResponse)
}

