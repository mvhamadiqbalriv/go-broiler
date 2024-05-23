package main

import (
	"mvhamadiqbalriv/belajar-golang-restful-api/app"
	"mvhamadiqbalriv/belajar-golang-restful-api/controller"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"mvhamadiqbalriv/belajar-golang-restful-api/repository"
	"mvhamadiqbalriv/belajar-golang-restful-api/service"
	"mvhamadiqbalriv/belajar-golang-restful-api/validator"
	"net/http"
)

func main() {

	db := app.NewDB()

	//with custom validator from dir validator
	validate := validator.NewCustomValidator()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

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