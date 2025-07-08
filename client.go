package financialmodelingprep

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

type restyDoer struct {
	client *resty.Client
}

func newRestyDoer() *restyDoer {
	c := resty.New()
	c.Debug = true
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
	return resp.RawResponse, nil
}

func MustClient(apikey string) ClientInterface {
	httpClientOption := WithHTTPClient(newRestyDoer())
	apiKeyProvider, err := securityprovider.NewSecurityProviderApiKey("query", "apikey", apikey)
	if err != nil {
		panic(err)
	}
	opts := []ClientOption{
		httpClientOption,
		WithRequestEditorFn(apiKeyProvider.Intercept),
	}
	c, err := NewClient("https://financialmodelingprep.com/stable", opts...)
	if err != nil {
		panic(err)
	}
	return c
}
