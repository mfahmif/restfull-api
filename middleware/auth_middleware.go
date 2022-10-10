package middleware

import (
	"net/http"
	"restful-api/helper"
	"restful-api/model/web"
)

type Auth_Middleware struct {
	Handler http.Handler
}

func New_Auth_Middleware(handler http.Handler) *Auth_Middleware {

	return &Auth_Middleware{Handler: handler}

}

func (auth *Auth_Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if "Value" == r.Header.Get("X-Api-Key") {
		auth.Handler.ServeHTTP(w, r)
		//oke
	} else {
		//error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteToResponseBody(w, webResponse)
	}

}
