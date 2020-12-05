package cache

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CacheTestSuite struct {
	suite.Suite
}

var MockRedis = NewMockRedis()

var testKey = "test_key"
var testVal = "test_val"

func (suite *CacheTestSuite) SetupTest() {
	MockRedis = NewMockRedis()
}

func (suite *CacheTestSuite) TestMockService_SetKeyStringRedis() {
	err := MockRedis.SetKeyStringRedis(testKey, testVal)
	suite.NoError(err)

	exists := MockRedis.IsKeyInRedis(testKey)
	suite.True(exists)

	err = MockRedis.SetKeyStringRedis(testKey, "new_val")
	suite.NoError(err)
}

func (suite *CacheTestSuite) TestMockService_IsKeyInRedis() {
	exists := MockRedis.IsKeyInRedis(testKey)
	suite.False(exists)

	err := MockRedis.SetKeyStringRedis(testKey, testVal)
	suite.NoError(err)

	exists = MockRedis.IsKeyInRedis(testKey)
	suite.True(exists)
}

func (suite *CacheTestSuite) TestMockService_ExpireKey() {
	err := MockRedis.SetKeyStringRedis(testKey, testVal)
	suite.NoError(err)

	MockRedis.ExpireKey(testKey)

	exists := MockRedis.IsKeyInRedis(testKey)
	suite.False(exists)
}

func (suite *CacheTestSuite) TestMockService_GetKeyInRedis() {
	err := MockRedis.SetKeyStringRedis(testKey, testVal)
	suite.NoError(err)

	val := MockRedis.GetKeyInRedis(testKey)
	suite.Equal(testVal, val)
}

// Run the Cache service test suite.
func TestHandlers(t *testing.T) {
	suite.Run(t, new(CacheTestSuite))
}
