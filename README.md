# go-financialmodelingprep

Offer an standardarized Go client for Financial Modeling Prep API.

### Install

```bash
go get github.com/zhoub/go-financialmodelingprep
```

### Usage

```bash
// Create client with API key.
c := MustClient(&ClientOptions{APIKey: apiKey})

// Call API to get data.
c.GetCompanyProfile(context.Background(), &ProfileGetParams{Symbol: "AAPL"})
```
