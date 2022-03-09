package paymentprovider

import (
	"ascenda_checkout_api_tests/services/payment_provider/provider_a"
	"ascenda_checkout_api_tests/services/response"
	"fmt"
)

func CheckOut(provider string, creditCardID string, amount float64, concurrency string) (res response.Response, err error) {
	switch provider {
	case provider_a.ProviderName:
		res, err = provider_a.CheckOut(creditCardID, amount, concurrency)
	default:
		err = fmt.Errorf("Invalid Provider Name")
	}

	return res, err
}
