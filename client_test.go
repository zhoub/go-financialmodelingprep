package financialmodelingprep

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type clientSuite struct {
	suite.Suite
	suite.SetupAllSuite

	c ClientInterface
}

func (r *clientSuite) SetupSuite() {
	apiKey := os.Getenv("FMP_API_KEY")
	r.c = MustClient(apiKey)
}

func (r *clientSuite) TestGetCompanyProfile() {
	if resp, err := r.c.ProfileGet(context.Background(), &ProfileGetParams{
		Symbol: "AAPL",
	}); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode)
	}
}

func (r *clientSuite) TestBatchQuoteShortGet() {
	if resp, err := r.c.BatchQuoteShortGet(context.Background(), &BatchQuoteShortGetParams{
		Symbols: "7201.T,7203.T",
	}); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode)
	}
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(clientSuite))
}
