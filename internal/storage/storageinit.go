package storage

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
	"go.uber.org/zap"
)

const migrationsDir = "C:\\Users\\frant\\go\\go1.21.0\\bin\\pkg\\mod\\github.com\\fraaaaaaaaaanc\\gofermart\\internal\\migrations"

type Storage struct {
	DB  *sql.DB
	log *zap.Logger
}

func NewStorage(storageDBAddress string, log *zap.Logger) (*Storage, error) {
	db, err := sql.Open("pgx", storageDBAddress)
	if err != nil {
		return nil, err
	}

	ctx, cansel := context.WithCancel(context.Background())
	defer cansel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	_, err = goose.GetDBVersion(db)
	if err != nil {
		return nil, err
	}

	err = goose.Up(db, migrationsDir)
	if err != nil && err != goose.ErrNoNextVersion {
		return nil, err
	}

	//_, err = db.ExecContext(ctx, `
	//DO $$
	//BEGIN
	//    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
	//        CREATE TABLE users (
	//            id SERIAL PRIMARY KEY,
	//            user_name VARCHAR(50) NOT NULL DEFAULT '' UNIQUE,
	//            password VARCHAR(60) NOT NULL DEFAULT ''
	//        );
	//    END IF;
	//
	//    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'orders') THEN
	//        CREATE TABLE orders (
	//            id SERIAL PRIMARY KEY,
	//            user_id INT REFERENCES users(id),
	//            order_number VARCHAR(32) NOT NULL DEFAULT '' UNIQUE,
	//            order_status VARCHAR(10) NOT NULL DEFAULT '',
	//            accrual DECIMAL NOT NULL DEFAULT 0,
	//            order_datetime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	//        );
	//	END IF;
	//
	//	IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'orders_info') THEN
	//        CREATE TABLE orders_info (
	//            id SERIAL PRIMARY KEY,
	//            order_id INT REFERENCES orders(id),
	//            description VARCHAR(128) NOT NULL DEFAULT '',
	//            price DECIMAL NOT NULL DEFAULT 0
	//        );
	//	END IF;
	//
	//	IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'balance') THEN
	//	   CREATE TABLE balance (
	//	       id SERIAL PRIMARY KEY,
	//	       user_id INT REFERENCES users(id),
	//	       user_balance DECIMAL NOT NULL DEFAULT 0 CHECK (user_balance >= 0),
	//	       withdrawn_balance DECIMAL NOT NULL DEFAULT 0
	//	    );
	//	END IF;
	//
	//	IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'history_balance') THEN
	//	   CREATE TABLE history_balance (
	//	       id SERIAL PRIMARY KEY,
	//	       user_id INT REFERENCES users(id),
	//	       order_number_unregister VARCHAR(32) NOT NULL DEFAULT '' UNIQUE,
	//	       withdrawn_sum DECIMAL NOT NULL DEFAULT 0 CHECK (withdrawn_sum >= 0),
	//	       withdrawn_datetime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	//	   );
	//    END IF;
	//
	//    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'order_accrual') THEN
	//	   CREATE TABLE order_accrual (
	//	       id SERIAL PRIMARY KEY,
	//	       user_id INT REFERENCES users(id),
	//	       order_id INT REFERENCES orders(id),
	//	       order_number VARCHAR(32) NOT NULL DEFAULT '' UNIQUE,
	//	       order_status_accrual VARCHAR(10) NOT NULL DEFAULT '',
	//	       accrual DECIMAL NOT NULL DEFAULT 0
	//	   );
	//    END IF;
	//END $$;
	//`)
	return &Storage{
		DB:  db,
		log: log,
	}, nil
}
