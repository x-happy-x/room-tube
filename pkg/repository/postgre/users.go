package postgre

import (
	"database/sql"
	"errors"
	"tube/pkg/model"
)

func (r *Repository) GetUserByID(id int64) (*model.User, error) {
	user := &model.User{}
	err := r.DB.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *Repository) CreateUser(user *model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *Repository) UpdateUser(user *model.User) error {
	_, err := r.DB.Exec("UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?", user.Name, user.Email, user.Password, user.UpdatedAt, user.ID)
	return err
}

func (r *Repository) DeleteUser(id int64) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
