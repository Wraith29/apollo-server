package data

func SaveUser(userId, username string) error {
	query := `INSERT INTO "users" ("user_id", "name") VALUES ($1, $2)`

	_, err := _conn.Exec(query, userId, username)

	return err
}
