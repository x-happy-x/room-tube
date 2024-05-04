package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"tube/pkg/api"
	"tube/pkg/model"
	"tube/pkg/service"
)

type UserHandler struct {
	handler http.Handler
	service service.Service
}

func (h *UserHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var user model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	h.service.User.
		// Получить данные авторизации из запроса
		username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Missing authorization data", http.StatusUnauthorized)
		return
	}

	// Проверить учетные данные пользователя
	user, err := api.userService.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Здесь вы можете установить cookie или другой механизм аутентификации

	fmt.Fprintf(w, "Hello, %s!", user.Name)
}
