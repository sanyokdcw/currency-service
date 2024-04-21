package currencyApi

type CurrencyClient interface {
	GetCurrencyRate(currencyCode string) (float64, error)
}

type API1Client struct{}
type API2Client struct{}

func (client *API1Client) GetCurrencyRate(currencyCode string) (float64, error) {
	return 1.0, nil
}

func (client *API2Client) GetCurrencyRate(currencyCode string) (float64, error) {
	return 1.2, nil
}
