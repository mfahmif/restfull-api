package exception

import (
	"net/http"
	"restful-api/helper"
	"restful-api/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writter http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writter, request, err) {
		return
	}

	if badRequest(writter, request, err) {
		return
	}

	internalServerError(writter, request, err)

}

func badRequest(writter http.ResponseWriter, request *http.Request, err interface{}) bool {

	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Status bad request",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writter, webResponse)

		return true
	} else {
		return false
	}

}

func notFoundError(writter http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writter.Header().Set("Content-Type", "application/json")
		writter.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Status not found error",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writter, webResponse)

		return true
	} else {
		return false
	}

}

func internalServerError(writter http.ResponseWriter, request *http.Request, err interface{}) {

	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(writter, webResponse)

}
