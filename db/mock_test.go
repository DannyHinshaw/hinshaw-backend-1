package db

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockDatabaseTestSuite struct {
	suite.Suite
}

func (suite *MockDatabaseTestSuite) TestMockService_ParseDB() {
	invalid := "DB_URL"
	err := DBMockService.ParseDB(invalid)
	suite.Error(err)

	valid := "postgres://user:pass@localhost:5432/db"
	err = DBMockService.ParseDB(valid)
	suite.NoError(err)

	c := DBMockService.Config
	suite.Equal(c.ConnConfig.ConnString(), valid)
}

func (suite *MockDatabaseTestSuite) TestMockService_Init() {
	err := DBMockService.Init()
	suite.NoError(err)
}

func (suite *MockDatabaseTestSuite) TestMockService_QueryAllCustomers() {
	ctx := context.Background()
	customers, err := DBMockService.QueryAllCustomers(ctx)
	suite.NoError(err)
	suite.Len(customers, 4)
}

// Run the MockDatabaseTestSuite test suite.
func TestHandlers(t *testing.T) {
	suite.Run(t, new(MockDatabaseTestSuite))
}
