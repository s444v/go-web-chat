package database

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// CheckPass проверяет пароль и возвращает список ролей пользователя.
func CheckPass(username, password string) (bool, []string, error) {
	query := `
		SELECT u.id, p.password_hash
		FROM users u
		JOIN user_passwords p ON u.id = p.id_user
		WHERE u.name = $1
	`
	var userID int
	var hashedPassword string

	err := DB.QueryRow(query, username).Scan(&userID, &hashedPassword)
	if err == sql.ErrNoRows {
		return false, nil, nil // пользователя нет
	} else if err != nil {
		return false, nil, err
	}

	// Сравниваем пароль
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false, nil, nil // пароль неверный
	}

	// Достаём роли пользователя
	rolesQuery := `
		SELECT r.role_name
		FROM user_roles ur
		JOIN roles r ON ur.id_role = r.id
		WHERE ur.id_user = $1
	`
	rows, err := DB.Query(rolesQuery, userID)
	if err != nil {
		return false, nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return true, nil, err
		}
		roles = append(roles, role)
	}

	return true, roles, nil
}
