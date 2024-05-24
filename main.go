package main

import (
	"mvhamadiqbalriv/belajar-golang-restful-api/app"
	"mvhamadiqbalriv/belajar-golang-restful-api/controller"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"mvhamadiqbalriv/belajar-golang-restful-api/repository"
	"mvhamadiqbalriv/belajar-golang-restful-api/service"
	"mvhamadiqbalriv/belajar-golang-restful-api/validator"
	"net/http"
	"os"
	"strings"
)

func main() {

	db := app.NewDB()

	//with custom validator from dir validator
	validate := validator.NewCustomValidator()
	mailService := service.NewMailService(
		os.Getenv("MAIL_FROM_ADDRESS"),
		os.Getenv("MAIL_HOST"),
		helper.StringToInt(os.Getenv("MAIL_PORT")),
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService, mailService)

	tokenBlacklistService := service.NewTokenBlacklistService()

	authController := controller.NewAuthController(userService, tokenBlacklistService)
	
	router := app.NewRouter(userController, authController)

	server := http.Server{
		Addr:    "localhost:3001",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

func intercept(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/") {
            http.NotFound(w, r)
            return
        }

        next.ServeHTTP(w, r)
    })
}