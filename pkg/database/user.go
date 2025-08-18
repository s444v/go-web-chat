package database

type Account struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func GetAccounts(limit int, search string) ([]*Account, error) {
	var users []*Account

	baseQuery := `
        SELECT * 
        FROM users
    `

	params := map[string]interface{}{
		"limit": limit,
	}

	if search != "" {
		baseQuery += `
            WHERE name ILIKE :search OR email ILIKE :search
        `
		params["search"] = "%" + search + "%"
	}

	baseQuery += `
        LIMIT :limit
    `

	stmt, err := DB.PrepareNamed(baseQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.Select(&users, params)
	if err != nil {
		return nil, err
	}

	if users == nil {
		users = make([]*Account, 0)
	}
	return users, nil
}
