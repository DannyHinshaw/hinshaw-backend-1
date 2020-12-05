package db

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"hinshaw-backend-1/schemas"
	"time"
)

type MockDatabase struct {
	Users        []*schemas.AppUser
	CreditScores []*schemas.CreditScore
	Customers    []*schemas.Customer
}

var MockDB = MockDatabase{
	Users:        []*schemas.AppUser{},
	CreditScores: []*schemas.CreditScore{},
	Customers:    []*schemas.Customer{},
}

// Handles creation of new test app user object.
func CreateNewUser(email, password string) (*schemas.AppUser, error) {
	hashPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	now := time.Now().String()
	newUser := &schemas.AppUser{
		Id:        uuid.NewV4().String(),
		Email:     email,
		Password:  hashPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return newUser, nil
}

// Handles retrieving a users id and password from db.
func (s *MockService) QueryUserAuth(email string, ctx context.Context) (*UserAuth, error) {
	for _, u := range MockDB.Users {
		if u.Email == email {
			userAuth := &UserAuth{
				UserId:   u.Id,
				Password: u.Password,
			}

			return userAuth, nil
		}
	}

	return nil, nil
}

// Handles querying all customers and their scores from postgres.
func (s *MockService) QueryAllCustomers(ctx context.Context) ([]CustomerScores, error) {
	return []CustomerScores{
		{
			FirstName:  "Kimberley",
			LastName:   "Rios",
			Equifax:    500,
			Experian:   550,
			TransUnion: 525,
		},
		{
			FirstName:  "Yee",
			LastName:   "Robinson",
			Equifax:    631,
			Experian:   617,
			TransUnion: 620,
		},
		{
			FirstName:  "Charles",
			LastName:   "Lainez",
			Equifax:    700,
			Experian:   770,
			TransUnion: 725,
		},
		{
			FirstName:  "James",
			LastName:   "Cannon",
			Equifax:    600,
			Experian:   615,
			TransUnion: 630,
		},
	}, nil
}

// Mock check if user already exists in MockDB.
func (s *MockService) QueryUserEmailExists(email string, ctx context.Context) (bool, error) {
	for _, u := range MockDB.Users {
		if u.Email == email {
			return true, nil
		}
	}

	return false, nil
}

// Mock add new user to db.
func (s *MockService) AddNewUser(email string, password string, ctx context.Context) error {
	newUser, err := CreateNewUser(email, password)
	if err != nil {
		return err
	}

	MockDB.Users = append(MockDB.Users, newUser)
	return nil
}
