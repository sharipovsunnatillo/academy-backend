package api

import (
	"encoding/json"
	"fmt"
	"github.com/sharipov/sunnatillo/academy-backend/pkg/dto"
	"net/http"
)

func UserMux() *http.ServeMux {
	mux := http.ServeMux{}
	mux.HandleFunc("POST /register", func(writer http.ResponseWriter, request *http.Request) {
		var createDto dto.UserCreateDto
		err := json.NewDecoder(request.Body).Decode(
			&createDto,
		)
		if err != nil {
			return
		}
		fmt.Println(createDto)
		writer.WriteHeader(http.StatusOK)
	})
	return &mux
}
