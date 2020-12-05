package middleware

import (
	"github.com/stretchr/testify/suite"
	td "hinshaw-backend-1/test"
	"testing"
)

type JWTTestSuite struct {
	suite.Suite
}

func (suite *JWTTestSuite) TestGenerateJWT() {
	jwtPayload, err := GenerateJWT(td.UserUUID)
	suite.NoError(err)
	suite.NotNil(jwtPayload)
	suite.NotNil(jwtPayload.AccessToken)
	suite.NotNil(jwtPayload.RefreshToken)
}

func (suite *JWTTestSuite) TestValidateJWT() {
	jwtPayload, err := GenerateJWT(td.UserUUID)
	suite.NoError(err)
	suite.NotNil(jwtPayload)
	suite.NotNil(jwtPayload.AccessToken)
	suite.NotNil(jwtPayload.RefreshToken)

	valid, err := ValidateJWT(jwtPayload.AccessToken)
	suite.NoError(err)
	suite.True(valid)
}

// Run the JWTTestSuite test suite.
func TestJWT(t *testing.T) {
	suite.Run(t, new(JWTTestSuite))
}
