package database

import "database/sql"

func GetOrCreateChat(user1, user2 int) (int, error) {
	var chatID int

	query := `
		SELECT c.id
		FROM chats c
		JOIN chat_members m1 ON m1.chat_id = c.id AND m1.user_id = $1
		JOIN chat_members m2 ON m2.chat_id = c.id AND m2.user_id = $2
		WHERE c.is_group = false
		LIMIT 1
	`
	err := DB.QueryRow(query, user1, user2).Scan(&chatID)
	if err == nil {
		return chatID, nil
	}
	if err != sql.ErrNoRows {
		return 0, err
	}

	tx, err := DB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	err = tx.QueryRow(
		`INSERT INTO chats (is_group) VALUES (false) RETURNING id`,
	).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	_, err = tx.Exec(
		`INSERT INTO chat_members (chat_id, user_id) VALUES ($1, $2), ($1, $3)`,
		chatID, user1, user2,
	)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return chatID, nil
}
