package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
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
var connectionRetries = 0
var maxConnectRetries = 5

// Parse the database url config to pgx pool.
func (s *Service) parse(dbURL string) error {
	var err error
	s.Config, err = pgxpool.ParseConfig(dbURL)
	s.Config.ConnConfig.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on",
	}
	s.Config.ConnConfig.PreferSimpleProtocol = true

	return err
}

// Initialize pgx pool for application use.
func (s *Service) init() error {
	var err error
	ctx := context.Background()
	s.Pool, err = pgxpool.ConnectConfig(ctx, s.Config)
	if connectionRetries >= maxConnectRetries {
		return err
	}

	// If successful
	if err == nil {

		// Seed the database with example data.
		if err := s.seed(); err != nil {
			log.Println("seed error::", err)
		}

		return nil
	}

	log.Println("pgxpool::err", err)
	message := "pgxpool connection failed:: sleeping %v seconds then retrying"
	log.Printf(message, connectionRetries)

	// Exponential backoff waiting for postgres...
	duration := time.Duration(connectionRetries)
	time.Sleep(duration * time.Second)
	connectionRetries++

	return s.init()
}

// Handles seeding the database with seed.go file.
func (s *Service) seed() error {
	path, err := filepath.Abs("seed.sql")
	if err != nil {
		log.Fatal("error retrieving seed.sql file path::", err)
		return err
	}

	bytes, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("error reading seed.sql file::", err)
		return err
	}

	ctx := context.Background()
	sql := string(bytes)
	_, err = s.Pool.Exec(ctx, sql)
	if err != nil {
		log.Println("error executing seed SQL::", err)
		return err
	}

	log.Println("database seeded successfully")
	return nil
}

// Util wrapper for parse to handle error.
func (s *Service) ParseDB(dbURL string) error {
	return s.parse(dbURL)
}

// Util wrapper to initialize pgx pool and seed db.
func (s *Service) Init() error {

	// Init pgx pool/config.
	if err := s.init(); err != nil {
		return err
	}

	return nil
}

// Util func to close pgx pool.
func (s *Service) Close() {
	s.Pool.Close()
}
