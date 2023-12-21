package middleware

import (
	"go-crud/helper"
	"go-crud/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (authMiddleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "RAHASIA" {
		authMiddleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
