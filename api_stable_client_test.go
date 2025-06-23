package gofmp

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type apiStableClientSuite struct {
	suite.Suite
}

func (r *apiStableClientSuite) TestGetSwagger() {
	s, err := GetSwagger()
	r.NotNil(s)
	r.NoError(err)
}

func TestApiStableClientSuite(t *testing.T) {
	suite.Run(t, &apiStableClientSuite{})
}
