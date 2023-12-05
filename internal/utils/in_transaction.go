package utils

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"gofermart/internal/logger"
	"time"
)

func InTransaction(db *sql.DB, reqCtx context.Context, f func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	ctx, cansel := context.WithTimeout(reqCtx, time.Second*1)
	defer cansel()

	if err := f(ctx, tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Error("error in the operation of tx.Rollback", zap.Error(rollbackErr))
		}
		return err
	}

	return tx.Commit()
}
