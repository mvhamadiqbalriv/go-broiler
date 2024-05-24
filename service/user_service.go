package service

import (
	"context"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/web/user_web"
	"mvhamadiqbalriv/belajar-golang-restful-api/pkg"
	"net/http"
)

type UserService interface {
	Create(ctx context.Context, request user_web.CreateRequest) user_web.Response
	Update(ctx context.Context, request user_web.UpdateRequest) user_web.Response
	Delete(ctx context.Context, userId int)
	FindByID(ctx context.Context, userId int) user_web.Response
	FindAll(ctx context.Context, r *http.Request) *pkg.PaginationImpl

	CreateProfilePicture(ctx context.Context, request user_web.CreateProfilePictureRequest) user_web.Response
	ChangePassword(ctx context.Context, request user_web.ChangePasswordRequest) user_web.Response

	Login(ctx context.Context, request user_web.LoginRequest) user_web.AuthResponse
}