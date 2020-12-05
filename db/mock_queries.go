package db

import "context"

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
