package api

import (
	"encoding/json"
	"github.com/sharipov/sunnatillo/academy-backend/internal/service"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/dto"
	"net/http"
)

type UserApi struct {
	userService *service.UserService
}

func NewUserApi(userService *service.UserService) *UserApi {
	return &UserApi{userService: userService}
}

func (userApi UserApi) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", func(writer http.ResponseWriter, request *http.Request) {
		var createDto dto.UserCreateDto
		if err := json.NewDecoder(request.Body).Decode(&createDto); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusCreated)
		register, err := userApi.userService.Register(createDto)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(writer).Encode(register); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	return mux
}

func (userApi UserApi) Register(mux *http.ServeMux) {
	mux.Handle("/api/users/v1/", http.StripPrefix("/api/users/v1", userApi.Routes()))
}
