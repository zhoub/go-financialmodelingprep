package financialmodelingprep

import (
	"bytes"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

type restyDoer struct {
	client *resty.Client
}

func newRestyDoer(debug bool) *restyDoer {
	c := resty.New()
	c.Debug = debug
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

type ClientOptions struct {
	APIKey string

	Debug bool
}

func MustClient(clientOptions *ClientOptions) *ClientWithResponses {
	httpClientOption := WithHTTPClient(newRestyDoer(clientOptions.Debug))
	apiKeyProvider, err := securityprovider.NewSecurityProviderApiKey("query", "apikey", clientOptions.APIKey)
	if err != nil {
		panic(err)
	}
	client, err := NewClientWithResponses("https://financialmodelingprep.com/stable",
		[]ClientOption{
			httpClientOption,
			WithRequestEditorFn(apiKeyProvider.Intercept),
		}...)
	if err != nil {
		panic(err)
	}
	return client
}
