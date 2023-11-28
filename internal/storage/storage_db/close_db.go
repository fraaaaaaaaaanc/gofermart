package storagedb

import (
	"go.uber.org/zap"
	"gofermart/internal/logger"
)

func (s *Storage) CloseDB() {
	if err := s.db.Close(); err != nil {
		logger.Error("error closing database connection", zap.Error(err))
	}
}
