package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IService interface {
	parse(dbURL string) error
	ParseDB(dbURL string) error
	init() error
	Init() error
	Close()
}

type Service struct {
	*pgxpool.Pool
	*pgxpool.Config
}

var DatabaseService = Service{}

// Parse the database url config to pgx pool.
func (s *Service) parse(dbURL string) error {
	var err error
	s.Config, err = pgxpool.ParseConfig(dbURL)
	s.Config.ConnConfig.RuntimeParams = map[string]string{
		"standard_conforming_string": "on",
	}
	s.Config.ConnConfig.PreferSimpleProtocol = true

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
func (s *Service) ParseDB(dbURL string) error {
	return s.parse(dbURL)
}

// Util wrapper to initialize pgx pool.
func (s *Service) Init() error {
	return s.init()
}

// Util func to close pgx pool.
func (s *Service) Close() {
	s.Pool.Close()
}
