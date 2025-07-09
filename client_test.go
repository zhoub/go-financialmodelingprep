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

	c ClientWithResponsesInterface
}

func (r *clientSuite) SetupSuite() {
	apiKey := os.Getenv("FMP_API_KEY")
	r.c = MustClient(&ClientConfig{
		APIKey: apiKey,
		Debug:  true,
	})
}

func (r *clientSuite) TestCompanyProfileGet() {
	symbol := "AAPL"
	if resp, err := r.c.ProfileGetWithResponse(context.Background(), &ProfileGetParams{
		Symbol: symbol,
	}); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode())

		pList := *resp.JSON200
		r.Len(pList, 1)
		r.Equal(symbol, pList[0].Symbol)
		r.Equal(1980, pList[0].IpoDate.Year())
	}
}

func (r *clientSuite) TestBatchQuoteShortGet() {
	if resp, err := r.c.BatchQuoteShortGetWithResponse(context.Background(), &BatchQuoteShortGetParams{
		Symbols: "7201.T,7203.T",
	}); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode())

		sqList := *resp.JSON200
		r.Len(sqList, 2)
	}
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(clientSuite))
}
