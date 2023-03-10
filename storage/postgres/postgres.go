package postgres

import (
	"context"
	"fmt"
	"rest_service/config"
	"rest_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db    *pgxpool.Pool
	phone storage.PhoneRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Phone() storage.PhoneRepoI {
	if s.phone == nil {
		s.phone = NewUserRepo(s.db)
	}
	return s.phone
}
