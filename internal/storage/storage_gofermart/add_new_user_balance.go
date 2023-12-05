package storagegofermart

import (
	"context"
	"database/sql"
)

func (s *Storage) AddNewUserBalance(ctx context.Context, tx *sql.Tx, userID int) error {
	_, err := tx.ExecContext(ctx,
		"INSERT INTO balance (user_id)"+
			"VALUES ($1)",
		userID)

	return err
}
