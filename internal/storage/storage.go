package storage

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(storageDBAddress string) (*Storage, error) {
	db, err := sql.Open("pgx", storageDBAddress)
	if err != nil {
		return nil, err
	}

	ctx, cansel := context.WithCancel(context.Background())
	defer cansel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, `
			DO $$ 
			BEGIN
				IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
					CREATE TABLE users (
						id SERIAL PRIMARY KEY, 
						UserName VARCHAR(50) NOT NULL DEFAULT '' UNIQUE, 
						Password VARCHAR(50) NOT NULL DEFAULT ''
					);
				END IF;
			END $$;
			`)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
