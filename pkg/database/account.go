package database

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func GetUserId(username string) (int, error) {
	var userId int
	err := DB.Get(&userId, "SELECT id FROM users WHERE name = $1", username)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func GetAccounts(userId, limit int, search string) ([]*Account, error) {
	var users []*Account

	query := `
		SELECT id, name, email
		FROM users
		WHERE id != :user_id
	`

	params := map[string]interface{}{
		"user_id": userId,
		"limit":   limit,
	}

	if search != "" {
		query += " AND (name ILIKE :search OR email ILIKE :search)"
		params["search"] = "%" + search + "%"
	}

	query += " ORDER BY id LIMIT :limit"

	rows, err := DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var acc Account
		if err := rows.StructScan(&acc); err != nil {
			return nil, err
		}
		users = append(users, &acc)
	}

	if users == nil {
		users = make([]*Account, 0)
	}

	return users, nil
}

func DeleteAccount(name string) error {

	query := `DELETE FROM users WHERE name = $1;`

	_, err := DB.Exec(query, name)
	if err != nil {
		return fmt.Errorf("failed to delete account: %w", err)
	}
	return nil
}

func AddAccount(name, password, email string) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	query := `
	WITH new_user AS (
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id
	),
	role_id AS (
		SELECT id AS rid FROM roles WHERE role_name = 'user'
	),
	ins_pass AS (
		INSERT INTO user_passwords (id_user, password_hash)
		SELECT id, $3
		FROM new_user
	)
	INSERT INTO user_roles (id_user, id_role)
	SELECT u.id, r.rid
	FROM new_user u, role_id r;
	`

	_, err = tx.Exec(query, name, email, string(hashedPass))
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" { // unique_violation
				return errors.New("account with this email already exists")
			}
		}
		return fmt.Errorf("failed to add account: %w", err)
	}

	return nil
}
