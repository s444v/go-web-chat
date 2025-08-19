package database

import (
	"fmt"
)

func AddMessege(senderUsername, recipientUsername, text string) error {
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

	query := `INSERT INTO messages (chat_id, sender_id, recipient_id, content)
VALUES (
    1,
    (SELECT id FROM users WHERE name = $1),
    (SELECT id FROM users WHERE name = $2),
    $3
);
`

	_, err = tx.Exec(query, senderUsername, recipientUsername, text)
	if err != nil {
		return fmt.Errorf("failed to save messege: %w", err)
	}

	return nil
}
