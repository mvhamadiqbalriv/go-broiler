package controller

import (
	"encoding/json"
	"fmt"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/web"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/web/user_web"
	"mvhamadiqbalriv/belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
	MailService service.MailService
}

func NewUserController(
	userService service.UserService,
	mailService service.MailService,
	) UserController {
	return &UserControllerImpl{
		UserService: userService,
		MailService: mailService,
	}
}

func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userCreateRequest := user_web.CreateRequest{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	userResponse := controller.UserService.Create(r.Context(), userCreateRequest)
	//send email
	body := fmt.Sprintf("Hi %s, <br> Welcome to our platform", userResponse.Name)
	err := controller.MailService.SendMail(userResponse.Email, "Welcome", body)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	userUpdateRequest := user_web.UpdateRequest{}
	helper.ReadFromRequestBody(r, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(r.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) FindByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindByID(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	usersResponse := controller.UserService.FindAll(r.Context(), r)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   usersResponse,
	}

	w.Header().Add("Content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *UserControllerImpl) ChangeProfilePicture(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	createProfilePictureRequest := user_web.CreateProfilePictureRequest{}
	helper.ReadFromRequestBody(r, &createProfilePictureRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	createProfilePictureRequest.Id = id

	userResponse := controller.UserService.CreateProfilePicture(r.Context(), createProfilePictureRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) ChangePassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	changePasswordRequest := user_web.ChangePasswordRequest{}
	helper.ReadFromRequestBody(r, &changePasswordRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	changePasswordRequest.Id = id

	userResponse := controller.UserService.ChangePassword(r.Context(), changePasswordRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
