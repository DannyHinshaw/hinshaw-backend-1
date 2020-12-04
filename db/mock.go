package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type MockService struct {
	*pgxpool.Pool
	*pgxpool.Config
}

var DBMockService = MockService{}

// Mock parse the database url config to pgx pool.
func (s *MockService) parse(dbURL string) error {
	log.Println("MockService::parse::dbURL::", dbURL)

	var err error
	s.Config, err = pgxpool.ParseConfig(dbURL)
	return err
}

// Mock initialize pgx pool for application use.
func (s *MockService) init() error {
	log.Println("MockService::init::")
	return nil
}

// Mock util wrapper for parse to handle error.
func (s *MockService) ParseDB(dbURL string) error {
	log.Println("MockService::ParseDB::dbURL::", dbURL)
	return s.parse(dbURL)
}

// Mock til wrapper to initialize pgx pool.
func (s *MockService) Init() error {
	log.Println("MockService::Init::")
	return s.init()
}

// Util func to close pgx pool.
func (s *MockService) Close() {
	log.Println("MockService::Close::")
	s.Pool.Close()
}
