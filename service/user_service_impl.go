package service

import (
	"context"
	"database/sql"
	"fmt"
	"mvhamadiqbalriv/belajar-golang-restful-api/exception"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/domain"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/web/user_web"
	"mvhamadiqbalriv/belajar-golang-restful-api/repository"
	"mvhamadiqbalriv/belajar-golang-restful-api/validator"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

const SecretKey = "secret"

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB			 *sql.DB
	validate	 *validator.CustomValidator
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.CustomValidator) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:				db,
		validate:		validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request user_web.CreateRequest) user_web.Response {
	err := service.validate.ValidateStruct(request)
	helper.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//check same email
	_, err = service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err == nil {
		panic(exception.NewDuplicateError("email already registered"))
	}

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	user = service.UserRepository.Create(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request user_web.UpdateRequest) user_web.Response {
	err := service.validate.ValidateStruct(request)
	helper.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByID(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Name = request.Name
	user.Email = request.Email

	//check same email except the current user
	userByEmail, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err == nil && user.ID != userByEmail.ID {
		panic(exception.NewDuplicateError("Email already registered"))
	}

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByID(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindByID(ctx context.Context, userId int) user_web.Response {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByID(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []user_web.Response {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUsersResponses(users)
}

func (service *UserServiceImpl) CreateProfilePicture(ctx context.Context, request user_web.CreateProfilePictureRequest) user_web.Response {
	err := service.validate.ValidateStruct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//check userId from request id and if null change to loggedUserId
	userId := request.Id
	if userId == 0 {
		userId = ctx.Value("loggedUserId").(int)
	}

	user, err := service.UserRepository.FindByID(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	//request is base 64 so we need to decode it and save it to the file
	profilePicture, err := helper.DecodeBase64(request.ProfilePicture)
	helper.PanicIfError(err)

	//save to file with dir assets/{randomFileName}
	profilePicturePath := fmt.Sprintf("assets/%s", helper.GenerateRandomString(10, "profile_picture"))
	profilePicturePath = fmt.Sprintf("%s.%s", profilePicturePath, "png")

	err = helper.SaveFile(profilePicturePath, profilePicture)
	helper.PanicIfError(err)

	//delete the old profile picture
	if user.ProfilePicture != "" {
		err = helper.DeleteFile(user.ProfilePicture)
		helper.PanicIfError(err)
	}

	user.ProfilePicture = profilePicturePath
	user = service.UserRepository.ChangeProfilePicture(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) ChangePassword(ctx context.Context, request user_web.ChangePasswordRequest) user_web.Response {
	err := service.validate.ValidateStruct(request)
	helper.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)


	//check userId from request id and if null change to loggedUserId
	userId := request.Id
	if userId == 0 {
		userId = ctx.Value("loggedUserId").(int)
	}

	fmt.Println("userId", userId)

	user, err := service.UserRepository.FindByID(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	cek := helper.ComparePasswords(user.Password, []byte(request.OldPassword))
	if !cek {
		panic(exception.NewBadRequestError("old password not match"))
	}

	user.Password = request.NewPassword
	user = service.UserRepository.ChangePassword(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request user_web.LoginRequest) user_web.AuthResponse {
	err := service.validate.ValidateStruct(request)
	helper.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	cek := helper.ComparePasswords(user.Password, []byte(request.Password))
	if !cek {
		panic(exception.NewUnauthorizedError("password not match"))
	}

	//generate token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
        Issuer:    strconv.Itoa(int(user.ID)), //issuer contains the ID of the user.
        ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //Adds time to the token i.e. 24 hours.
    })

    token, err := claims.SignedString([]byte(SecretKey))
	helper.PanicIfError(err)
	
	return helper.ToUserResponseAuth(user, token)
}