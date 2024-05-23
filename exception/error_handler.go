package exception

import (
	"fmt"
	"mvhamadiqbalriv/belajar-golang-restful-api/helper"
	"mvhamadiqbalriv/belajar-golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	// Open log file
	file, err := helper.OpenLogFile()
	if err != nil {
		fmt.Printf("Error opening log file: %s\n", err)
		// Exit or handle the error as needed
		return
	}

	// Initialize logger
	logger = logrus.New()
	logger.Out = file
	logger.Formatter = &logrus.JSONFormatter{}
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if unauthorizedError(writer, request, err) {
		logger.Info(err)
		return
	}

	if notFoundError(writer, request, err) {
		logger.Info(err)
		return
	}

	if validationErrors(writer, request, err) {
		logger.Info(err)
		return
	}

	if badRequestError(writer, request, err) {
		logger.Info(err)
		return
	}

	if duplicateError(writer, request, err) {
		logger.Info(err)
		return
	}

	logger.Error(err)
	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	validationErr, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		fieldErrors := make(map[string]string)
		for _, e := range validationErr {
			switch e.Tag() {
			case "required":
				fieldErrors[e.Field()] = "Field is required"
			case "confirmNewPassword":
				fieldErrors[e.Field()] = "Confirm new password must match with new password"
			default:
				fieldErrors[e.Field()] = e.Error()
			}
		}

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   fieldErrors,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	}else{
		return false
	}
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func duplicateError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(DuplicateError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusConflict)

		webResponse := web.WebResponse{
			Code:   http.StatusConflict,
			Status: "DUPLICATE",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}