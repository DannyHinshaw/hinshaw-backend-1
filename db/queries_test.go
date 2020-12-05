package db

import (
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type QueriesTestSuite struct {
	suite.Suite
}

func (suite *QueriesTestSuite) TestHashPassword() {
	password := "1234567890"
	hashedPassword, err := HashPassword(password)
	suite.NoError(err)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	suite.NoError(err)
}

// Run the Handlers test suite.
func TestQueries(t *testing.T) {
	suite.Run(t, new(QueriesTestSuite))
}
