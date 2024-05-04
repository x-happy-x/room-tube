package api

import (
	"fmt"
	"net/http"
)

type UserHandler struct {
}

func (api *roomTubeAPI) Authorize(w http.ResponseWriter, r *http.Request) {
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

func (api *roomTubeAPI) Register(w http.ResponseWriter, r *http.Request) {
	// Получить данные регистрации из запроса
	// Здесь вы можете использовать пакет `encoding/json` для декодирования JSON-запроса в структуру User

	// Создать нового пользователя
	err := api.userService.CreateUser(&user)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User %s created successfully", user.Name)
}

func (api *roomTubeAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	// Получить ID пользователя из параметров запроса
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	// Получить пользователя по ID
	user, err := api.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Здесь вы можете использовать пакет `encoding/json` для кодирования пользователя в JSON

	fmt.Fprintf(w, "User: %s", user.Name)
}
