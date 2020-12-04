package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	*pgxpool.Pool
	*pgxpool.Config
}

var DatabaseService = Service{}

// Parse the database url config to pgx pool.
func (s *Service) parse(dbURL string) error {
	var err error
	s.Config, err = pgxpool.ParseConfig(dbURL)
	return err
}

// Initialize pgx pool for application use.
func (s *Service) init() error {
	var err error
	ctx := context.Background()
	s.Pool, err = pgxpool.ConnectConfig(ctx, s.Config)
	return err
}

// Util wrapper for parse to handle error.
func ParseDB(dbURL string) error {
	return DatabaseService.parse(dbURL)
}

// Util wrapper to initialize pgx pool.
func Init() error {
	return DatabaseService.init()
}

// Util func to close pgx pool.
func Close() {
	DatabaseService.Pool.Close()
}
