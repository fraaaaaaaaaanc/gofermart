package storagedb

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"gofermart/internal/logger"
)

func (s *Storage) inTransaction(ctx context.Context, db *sql.DB, f func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := f(ctx, tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Error("error in the operation of tx.Rollback", zap.Error(rollbackErr))
		}
		return err
	}

	return tx.Commit()
}
