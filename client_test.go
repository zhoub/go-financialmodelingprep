package financialmodelingprep

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type clientSuite struct {
	suite.Suite
	suite.SetupAllSuite

	c *ClientWithResponses
}

func (r *clientSuite) SetupSuite() {
	apiKey := os.Getenv("FMP_API_KEY")
	r.c = MustClient(&ClientConfig{
		APIKey: apiKey,
		Debug:  true,
	})
}

func (r *clientSuite) TestAvailableExchanges() {
	if resp, err := r.c.AvailableExchangesGetWithResponse(context.Background()); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode())
		r.NotNil(resp.JSON200)

		xchgs := *resp.JSON200
		r.NotEmpty(xchgs)
	}
}

func (r *clientSuite) TestStockList() {
	if resp, err := r.c.StockListGetWithResponse(context.Background()); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode())
		r.NotNil(resp.JSON200)

		csList := *resp.JSON200
		r.NotEmpty(csList)
	}
}

func (r *clientSuite) TestProfileAAPL() {
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
		r.Contains(pList[0].IpoDate, "1980")
	}
}

func (r *clientSuite) TestProfileEQV() {
	const symbol = "EQV"
	params := map[string]interface{}{
		"symbol": symbol,
	}
	if resp, err := Get(context.Background(), r.c, ProfileGetOperationPath, params); err != nil {
		r.NoError(err)
	} else {
		r.NoError(err)
		r.Equal(http.StatusOK, resp.StatusCode)

		var profiles []CompanyProfile
		err = json.NewDecoder(resp.Body).Decode(&profiles)
		r.NoError(err)
		r.NotEmpty(profiles)

		// Validate symbol.
		r.Equal(symbol, profiles[0].Symbol)

		// The field "ipoDate" is empty string.
		r.Empty(profiles[0].IpoDate)
	}
}

var batchSymbos = []string{
	"1329.T",
	"7201.T",
	"7203.T",
	"7733.T",
	"AAPL",
	"AMZN",
	"MSFT",
	"PFE",
	"QQQ",
	"SPY",
	"TQQQ",
	"USDAUD",
	"USDCNY",
	"USDEUR",
	"USDJPY",
	"USDZAR",
}

func (r *clientSuite) TestBatchQuote() {
	if resp, err := r.c.BatchQuoteGetWithResponse(context.Background(), &BatchQuoteGetParams{
		Symbols: strings.Join(batchSymbos, ","),
	}); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode())

		sqList := *resp.JSON200
		r.Len(sqList, len(batchSymbos))
	}
}

func (r *clientSuite) TestBatchQuoteShort() {
	if resp, err := r.c.BatchQuoteShortGetWithResponse(context.Background(), &BatchQuoteShortGetParams{
		Symbols: strings.Join(batchSymbos, ","),
	}); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode())

		sqList := *resp.JSON200
		r.Len(sqList, len(batchSymbos))
	}
}

func (r *clientSuite) TestSharesFloatAMZN() {
	const symbol = "AMZN"
	params := map[string]interface{}{
		"symbol": symbol,
	}
	if resp, err := Get(context.Background(), r.c, SharesFloatGetOperationPath, params); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode)

		var csfList []CompanySharesFloat
		err = json.NewDecoder(resp.Body).Decode(&csfList)
		r.NoError(err)

		r.Len(csfList, 1)
		r.Equal(symbol, csfList[0].Symbol)
	}
}

func (r *clientSuite) TestSearchSymbolAMZN() {
	queries := map[string]interface{}{"query": "AMZN", "limit": 1}
	if resp, err := Get(context.Background(), r.c, SearchSymbolGetOperationPath, queries); err != nil {
		r.NoError(err)
	} else {
		r.Equal(http.StatusOK, resp.StatusCode)

		var csfList []CompanySharesFloat
		err = json.NewDecoder(resp.Body).Decode(&csfList)
		r.NoError(err)

		r.Len(csfList, 1)
		r.Equal(queries["query"], csfList[0].Symbol)
	}
}

func (r *clientSuite) TestBalanceSheetStatementAAPL() {
	const symbol = "AAPL"
	params := map[string]interface{}{
		"symbol": symbol,
		"period": "FY",
		"limit":  1,
	}
	if resp, err := Get(context.Background(), r.c, BalanceSheetStatementGetOperationPath, params); err != nil {
		r.NoError(err)
	} else {
		r.NoError(err)
		r.Equal(http.StatusOK, resp.StatusCode)

		var bsList []BalanceSheetStatement
		err = json.NewDecoder(resp.Body).Decode(&bsList)
		r.NoError(err)
		r.Len(bsList, 1)
		r.Equal(symbol, bsList[0].Symbol)
	}
}

func (r *clientSuite) TestBalanceSheetStatementSAIMC() {
	const symbol = "SAI.MC"
	params := map[string]interface{}{
		"symbol": symbol,
		"period": FY,
		"limit":  1,
	}
	if resp, err := Get(context.Background(), r.c, BalanceSheetStatementGetOperationPath, params); err != nil {
		r.NoError(err)
	} else {
		r.NoError(err)
		r.Equal(http.StatusOK, resp.StatusCode)

		var bssList []BalanceSheetStatement
		err = json.NewDecoder(resp.Body).Decode(&bssList)
		r.NoError(err)

		r.Len(bssList, 1)
		r.Equal(symbol, bssList[0].Symbol)
	}
}

func (r *clientSuite) TestIncomeStatementSAIMC() {
	const symbol = "SAI.MC"
	params := map[string]interface{}{
		"symbol": symbol,
		"period": FY,
		"limit":  2,
	}
	if resp, err := Get(context.Background(), r.c, IncomeStatementGetOperationPath, params); err != nil {
		r.NoError(err)
	} else {
		r.NoError(err)
		r.Equal(http.StatusOK, resp.StatusCode)

		var isList []IncomeStatement
		err = json.NewDecoder(resp.Body).Decode(&isList)
		r.NoError(err)

		r.Len(isList, params["limit"].(int))
		r.Equal(symbol, isList[0].Symbol)
	}
}

func (r *clientSuite) TestKeyMetricsNMG() {
	const symbol = "NMG"
	params := map[string]interface{}{
		"symbol": symbol,
		"period": FY,
		"limit":  2,
	}
	if resp, err := Get(context.Background(), r.c, KeyMetricsGetOperationPath, params); err != nil {
		r.NoError(err)
	} else {
		r.NoError(err)
		r.Equal(http.StatusOK, resp.StatusCode)

		var kmList []KeyMetrics
		err = json.NewDecoder(resp.Body).Decode(&kmList)
		r.NoError(err)

		r.Len(kmList, params["limit"].(int))
		for _, km := range kmList {
			r.Equal(symbol, km.Symbol)
		}
		r.Equal(time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC), kmList[0].Date.Time)
	}
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(clientSuite))
}
