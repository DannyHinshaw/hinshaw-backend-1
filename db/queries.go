package db

import "context"

type CustomerScores struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Equifax    int16  `json:"equifax"`
	Experian   int16  `json:"experian"`
	TransUnion int16  `json:"trans_union"`
}

// Handles querying all customers and their scores from postgres.
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
