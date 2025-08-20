package database

import (
	"fmt"
)

func AddMessage(chat_id int, sender_id int, text string) error {
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
    $1,
    $2,
    (SELECT user_id FROM chat_members WHERE chat_id = $1 and user_id != $2),
    $3
);
`

	_, err = tx.Exec(query, chat_id, sender_id, text)
	if err != nil {
		return fmt.Errorf("failed to save messege: %w", err)
	}

	return nil
}

func GetMessages(chat_id, limit int) ([]*Message, error) {
	rows, err := DB.Query(`
SELECT *
FROM (
    SELECT 
        m.id,
        m.chat_id,
        u.name AS sender,
        r.name AS receiver,
        m.content,
        m.created_at,
        m.is_read
    FROM messages m
    JOIN users u ON m.sender_id = u.id
    JOIN users r ON m.recipient_id = r.id
    WHERE m.chat_id = $1
    ORDER BY m.created_at DESC
    LIMIT $2
) sub
ORDER BY created_at ASC;
	`, chat_id, limit)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var messages []*Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.ChatID, &msg.Sender, &msg.Receiver, &msg.Text, &msg.CreatedAt, &msg.IsRead); err != nil {
			fmt.Println(err)
			return nil, err
		}
		messages = append(messages, &msg)
	}
	return messages, nil
}
