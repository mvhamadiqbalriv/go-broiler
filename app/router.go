package app

import (
	"mvhamadiqbalriv/belajar-golang-restful-api/controller"
	"mvhamadiqbalriv/belajar-golang-restful-api/exception"
	"mvhamadiqbalriv/belajar-golang-restful-api/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
		userController controller.UserController,
		authController controller.AuthController,
	) *httprouter.Router {
    router := httprouter.New()

	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/logout", middleware.AuthenticateMiddleware(authController.Logout))

	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:userId", userController.FindByID)

	router.POST("/api/users", middleware.AuthenticateMiddleware(userController.Create))
	router.PUT("/api/users/:userId", middleware.AuthenticateMiddleware(userController.Update))
	router.DELETE("/api/users/:userId", middleware.AuthenticateMiddleware(userController.Delete))
	router.PUT("/api/users/:userId/profile-picture", middleware.AuthenticateMiddleware(userController.ChangeProfilePicture))
	router.PUT("/api/users/:userId/change-password", middleware.AuthenticateMiddleware(userController.ChangePassword))

    // Panic handler
    router.PanicHandler = exception.ErrorHandler

    return router
}