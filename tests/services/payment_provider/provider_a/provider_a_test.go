package provider_a

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ascenda_checkout_api_tests/services/payment_provider/provider_a"
	"ascenda_checkout_api_tests/services/response"
)

func TestSuccessedCheckOut(t *testing.T) {
	ass := assert.New(t)
	creditCardID := "4000000000009995"
	amount := 999.00
	currency := "usd"

	res, err := provider_a.CheckOut(creditCardID, amount, currency)

	ass.Nil(err)
	ass.Equal(res.Status, response.StatusSuccessed)
	ass.Empty(res.FailureReason)
	ass.NotNil(res.Outcome)
}
