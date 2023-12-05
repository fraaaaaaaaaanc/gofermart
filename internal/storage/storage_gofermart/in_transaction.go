package storagegofermart

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
	"gofermart/internal/logger"
)

func (s *Storage) InTransaction(parentsCtx context.Context, f func(ctx context.Context, tx *sql.Tx) error) error {
	ctx, cansel := context.WithTimeout(parentsCtx, durationWorkCtx)
	defer cansel()
	tx, err := s.db.Begin()
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
