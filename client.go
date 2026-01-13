package financialmodelingprep

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

type restyDoer struct {
	client *resty.Client
}

func newRestyDoer(debug bool) *restyDoer {
	// Create resty clinet with retry.
	c := resty.New()

	c.Debug = debug

	c.RetryCount = 100
	c.RetryWaitTime = 1 * time.Minute

	return &restyDoer{
		client: c,
	}
}

func (r *restyDoer) Do(req *http.Request) (*http.Response, error) {
	// Convert http.Request to resty.Request
	restyReq := r.client.R().
		SetContext(req.Context())

	// Prepare URL.
	url := req.URL.String()

	// Send request.
	resp, err := restyReq.Execute(req.Method, url)
	if err != nil {
		return nil, err
	}
	rawResp := resp.RawResponse
	rawResp.Status = resp.Status()
	rawResp.StatusCode = resp.StatusCode()
	rawResp.Header = resp.Header()
	rawResp.Body = io.NopCloser(bytes.NewBuffer(resp.Body()))
	rawResp.ContentLength = resp.Size()
	return rawResp, nil
}

type ClientConfig struct {
	APIKey string

	Endpoint string

	Debug bool
}

func MustClient(cfg *ClientConfig) *ClientWithResponses {
	// Prepare server URL.
	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}
	server := swagger.Servers[0].URL
	if len(cfg.Endpoint) > 0 {
		server = cfg.Endpoint
	}

	// Prepare options.
	httpClientOption := WithHTTPClient(newRestyDoer(cfg.Debug))
	apiKeyProvider, err := securityprovider.NewSecurityProviderApiKey("query", "apikey", cfg.APIKey)
	if err != nil {
		panic(err)
	}
	clientOptions := []ClientOption{
		httpClientOption,
		WithRequestEditorFn(apiKeyProvider.Intercept),
	}

	// Return client.
	client, err := NewClientWithResponses(server, clientOptions...)
	if err != nil {
		panic(err)
	}
	return client
}

// Get executes an API operation using the provided client, context, operation path, and parameters.
func Get(ctx context.Context, c *ClientWithResponses, path OperationPath, params map[string]interface{}) (*http.Response, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	switch path {
	case AnalystEstimatesGetOperationPath:
		var p AnalystEstimatesGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.AnalystEstimatesGet(ctx, &p)
	case BalanceSheetStatementGetOperationPath:
		var p BalanceSheetStatementGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.BalanceSheetStatementGet(ctx, &p)
	case BalanceSheetStatementTTMGetOperationPath:
		var p BalanceSheetStatementTTMGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.BalanceSheetStatementTTMGet(ctx, &p)
	case BatchQuoteGetOperationPath:
		var p BatchQuoteGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.BatchQuoteGet(ctx, &p)
	case BatchQuoteShortGetOperationPath:
		var p BatchQuoteShortGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.BatchQuoteShortGet(ctx, &p)
	case CashFlowStatementGetOperationPath:
		var p CashFlowStatementGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.CashFlowStatementGet(ctx, &p)
	case CashFlowStatementTTMGetOperationPath:
		var p CashFlowStatementTTMGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.CashFlowStatementTTMGet(ctx, &p)
	case DelistedCompaniesOperationPath:
		var p DelistedCompaniesParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.DelistedCompanies(ctx, &p)
	case EconomicCalendarGetOperationPath:
		var p EconomicCalendarGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.EconomicCalendarGet(ctx, &p)
	case EsgDisclosuresGetOperationPath:
		var p EsgDisclosuresGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.EsgDisclosuresGet(ctx, &p)
	case EsgRatingsGetOperationPath:
		var p EsgRatingsGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.EsgRatingsGet(ctx, &p)
	case GradesLatestNewsGetOperationPath:
		var p GradesLatestNewsGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.GradesLatestNewsGet(ctx, &p)
	case HistoricalChart15MinGetOperationPath:
		var p HistoricalChart15MinGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.HistoricalChart15MinGet(ctx, &p)
	case HistoricalPriceEodFullGetOperationPath:
		var p HistoricalPriceEodFullGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.HistoricalPriceEodFullGet(ctx, &p)
	case HistoricalPriceEodLightGetOperationPath:
		var p HistoricalPriceEodLightGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.HistoricalPriceEodLightGet(ctx, &p)
	case IncomeStatementGetOperationPath:
		var p IncomeStatementGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.IncomeStatementGet(ctx, &p)
	case IncomeStatementTTMGetOperationPath:
		var p IncomeStatementTTMGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.IncomeStatementTTMGet(ctx, &p)
	case InsiderTradingLatestGetOperationPath:
		var p InsiderTradingLatestGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.InsiderTradingLatestGet(ctx, &p)
	case InsiderTradingSearchGetOperationPath:
		var p InsiderTradingSearchGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.InsiderTradingSearchGet(ctx, &p)
	case KeyMetricsGetOperationPath:
		var p KeyMetricsGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.KeyMetricsGet(ctx, &p)
	case MarketCapitalizationGetOperationPath:
		var p MarketCapitalizationGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.MarketCapitalizationGet(ctx, &p)
	case MarketCapitalizationBatchGetOperationPath:
		var p MarketCapitalizationBatchGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.MarketCapitalizationBatchGet(ctx, &p)
	case NewsGeneralLatestGetOperationPath:
		var p NewsGeneralLatestGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.NewsGeneralLatestGet(ctx, &p)
	case ProfileGetOperationPath:
		var p ProfileGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.ProfileGet(ctx, &p)
	case QuoteGetOperationPath:
		var p QuoteGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.QuoteGet(ctx, &p)
	case QuoteShortGetOperationPath:
		var p QuoteShortGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.QuoteShortGet(ctx, &p)
	case RatingsSnapshotGetOperationPath:
		var p RatingsSnapshotGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.RatingsSnapshotGet(ctx, &p)
	case RatiosGetOperationPath:
		var p RatiosGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.RatiosGet(ctx, &p)
	case RevenueGeographicSegmentationGetOperationPath:
		var p RevenueGeographicSegmentationGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.RevenueGeographicSegmentationGet(ctx, &p)
	case RevenueProductSegmentationGetOperationPath:
		var p RevenueProductSegmentationGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.RevenueProductSegmentationGet(ctx, &p)
	case SearchNameGetOperationPath:
		var p SearchNameGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.SearchNameGet(ctx, &p)
	case SearchSymbolGetOperationPath:
		var p SearchSymbolGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.SearchSymbolGet(ctx, &p)
	case SharesFloatGetOperationPath:
		var p SharesFloatGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.SharesFloatGet(ctx, &p)
	case TreasuryRatesGetOperationPath:
		var p TreasuryRatesGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.TreasuryRatesGet(ctx, &p)
	case AvailableExchangesGetOperationPath:
		resp, err = c.AvailableExchangesGet(ctx)
	case DividendsGetOperationPath:
		var p DividendsGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.DividendsGet(ctx, &p)
	case DividendsCalendarGetOperationPath:
		var p DividendsCalendarGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.DividendsCalendarGet(ctx, &p)
	case EarningsGetOperationPath:
		var p EarningsGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.EarningsGet(ctx, &p)
	case GetEarningsCalendarOperationPath:
		var p GetEarningsCalendarParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.GetEarningsCalendar(ctx, &p)
	case EnterpriseValueGetOperationPath:
		var p EnterpriseValueGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.EnterpriseValueGet(ctx, &p)
	case KeyMetricsTTMGetOperationPath:
		var p KeyMetricsTTMGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.KeyMetricsTTMGet(ctx, &p)
	case RatiosTTMGetOperationPath:
		var p RatiosTTMGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.RatiosTTMGet(ctx, &p)
	case GetSplitsOperationPath:
		var p GetSplitsParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.GetSplits(ctx, &p)
	case GetSplitsCalendarOperationPath:
		var p GetSplitsCalendarParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.GetSplitsCalendar(ctx, &p)
	case StockListGetOperationPath:
		resp, err = c.StockListGet(ctx)
	case NewsStockLatestGetOperationPath:
		var p NewsStockLatestGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.NewsStockLatestGet(ctx, &p)
	case TechnicalIndicatorsRsiGetOperationPath:
		var p TechnicalIndicatorsRsiGetParams
		if err := json.Unmarshal(paramsJSON, &p); err != nil {
			return nil, err
		}
		resp, err = c.TechnicalIndicatorsRsiGet(ctx, &p)
	default:
		return nil, fmt.Errorf("%s", string(path))
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
