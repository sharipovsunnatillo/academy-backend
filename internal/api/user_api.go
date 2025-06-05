package api

import (
	"github.com/sharipov/sunnatillo/academy-backend/internal/service"
	"net/http"
)

type UserApi struct {
	userService *service.UserService
}

func NewUserApi(userService *service.UserService) *UserApi {
	return &UserApi{userService: userService}
}

func (userApi UserApi) UserMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusCreated)
	})
	return mux
}
