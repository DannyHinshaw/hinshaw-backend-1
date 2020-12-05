package db

import (
	"context"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type CustomerScores struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Equifax    int16  `json:"equifax"`
	Experian   int16  `json:"experian"`
	TransUnion int16  `json:"trans_union"`
}

type UserAuth struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

// Util function to hash/salt a password for storage.
func HashPassword(password string) (string, error) {
	bytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println("error occurred while hashing password::", err)
		return "", err
	}

	return string(hashedPassword), nil
}

// Handles checking if a user already exists by email address.
func (s *Service) QueryUserEmailExists(email string, ctx context.Context) (bool, error) {
	var id string
	q := "SELECT id FROM app_users where email=$1;"
	err := s.Pool.QueryRow(ctx, q, email).Scan(&id)
	if err == nil {
		return true, nil
	}

	if id == "" {
		return false, nil
	}

	return true, nil
}

// Handles retrieving a users id and password from db..
func (s *Service) QueryUserAuth(email string, ctx context.Context) (*UserAuth, error) {
	var userAuth UserAuth
	q := "SELECT id, password FROM app_users where email=$1;"
	err := s.Pool.QueryRow(ctx, q, email).Scan(&userAuth.UserId, &userAuth.Password)
	if err != nil || userAuth.Password == "" {
		log.Printf("error querying password for %s:: %s", email, err)
		return nil, err
	}

	return &userAuth, nil
}

// Handles retrieving all customers from db.
func (s *Service) QueryAllCustomers(ctx context.Context) ([]CustomerScores, error) {
	q := "SELECT first_name, last_name, equifax, experian, trans_union FROM customers NATURAL JOIN credit_scores;"
	rows, err := s.Pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []CustomerScores
	for rows.Next() {

		var c CustomerScores
		targets := []interface{}{
			&c.FirstName, &c.LastName, &c.Equifax,
			&c.Experian, &c.TransUnion,
		}

		err = rows.Scan(targets...)
		if err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

// Handles inserting new users into db.
func (s *Service) AddNewUser(email string, password string, ctx context.Context) error {

	// Create user id
	userId := uuid.NewV4()

	// Hash/salt the password for db.
	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Println("error occurred while hashing password::", err)
		return err
	}

	q := "INSERT INTO app_users (id, email, password) VALUES ($1, $2, $3);"
	s.Pool.QueryRow(ctx, q, userId, email, hashedPassword)

	return nil
}
