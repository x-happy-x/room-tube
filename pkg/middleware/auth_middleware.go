package middleware

import (
	"context"
	"log"
	"net/http"
	"tube/pkg/model"
	"tube/pkg/utils"
)

type key int

const UserKey key = 0

func (m *Struct) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("AuthMiddleware: начало обработки запроса")

		// Получение токена из запроса
		token := r.Header.Get("Authorization")
		if token == "" { // Токен пуст
			log.Println("AuthMiddleware: токен отсутствует")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Валидация токена
		data, err := m.services.Auth.ValidateToken(token)
		if err != nil { // Токен неверный
			log.Printf("AuthMiddleware: ошибка при валидации токена: %v", err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Получение пользователя из токена
		var user model.User
		err = utils.ReMarshal(data["user"], &user)
		if err != nil {
			log.Printf("AuthMiddleware: ошибка при преобразовании пользователя: %v", err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Добавление в контекст
		ctx := context.WithValue(r.Context(), UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))

		log.Println("AuthMiddleware: завершение обработки запроса")
	})
}
