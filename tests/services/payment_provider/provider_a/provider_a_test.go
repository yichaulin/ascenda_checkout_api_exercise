package provider_a

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ascenda_checkout_api_tests/services/payment_provider/provider_a"
	"ascenda_checkout_api_tests/services/response"
)

func TestSuccessedCheckOut(t *testing.T) {
	ass := assert.New(t)
	amount := 999.00
	currency := "usd"
	creditCardID := "4242424242424242"

	res, err := provider_a.CheckOut(creditCardID, amount, currency)

	ass.Nil(err)
	ass.Equal(res.Status, response.StatusSuccessful)
	ass.Nil(res.FailureReason)
	ass.NotNil(res.Outcome)
}

func TestFailureCheckOut(t *testing.T) {
	ass := assert.New(t)
	amount := 999.00
	currency := "usd"

	tests := []struct {
		creditCardID          string
		expectedStatus        string
		expectedFailureReason string
	}{
		{"4000000000009979", response.StatusFailure, response.FailureReasonStolenCard},
		{"4000000000009995", response.StatusFailure, response.FailureReasonInsufficientFunds},
		{"4000000000009987", response.StatusFailure, response.FailureReasonLostCard},
		{"40000007600000", response.StatusFailure, response.FailureReasonInvalidCreditCard},
		{"4100000000000019", response.StatusFailure, response.FailureReasonBlocked},
	}

	for _, tc := range tests {
		res, err := provider_a.CheckOut(tc.creditCardID, amount, currency)

		ass.Nil(err)
		ass.Equal(res.Status, tc.expectedStatus)
		ass.Equal(*res.FailureReason, tc.expectedFailureReason)
		ass.NotNil(res.Outcome)
	}
}
