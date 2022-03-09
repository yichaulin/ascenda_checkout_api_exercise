package provider_a

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"

	"ascenda_checkout_api_tests/services/response"
)

const (
	APIUrl       = "http://5df9c4eee9f79e0014b6b2eb.mockapi.io/charge"
	ProviderName = "ProviderA"

	RespStatusSuccessed = "succeeded"
	RespStatusFailed    = "failed"

	RespOutcomeTypeIssuerDeclined = "issuer_declined"
	RespOutcomeTypeBlocked        = "blocked"

	RespOutcomeDeclinedReasonStolenCard        = "stolen_card"
	RespOutcomeDeclinedReasonLostCard          = "lost_card"
	RespOutcomeDeclinedReasonInsufficientFunds = "insufficient_funds"
	RespOutcomeDeclinedReasonHighestRiskLevel  = "highest_risk_level"

	RespErrorTypeInvalidReq       = "invalid_request_error"
	RespErrorInvalidCreditCardMsg = "Invalid Primary Account Number provided"
)

type checkOutRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type checkOutResponse struct {
	Status      string `json:"status"`
	Outcome     checkoutOutcome
	ID          string   `json:"id"`
	Object      string   `json:"object"`
	Amount      float64  `json:"amount"`
	Created     uint64   `json:"created"`
	Currency    string   `json:"currency"`
	Customer    string   `json:"customer"`
	Description string   `json:"description"`
	Error       reqError `json:"error"`
}

type checkoutOutcome struct {
	NetworkStatus string `json:"network_status"`
	Reason        string `json:"reason"`
	RiskLevel     string `json:"risk_level"`
	RiskScore     int64  `json:"risk_score"`
	SellerMessage string `json:"seller_message"`
	Type          string `json:"type"`
}

type reqError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func CheckOut(creditCardID string, amount float64, currency string) (res response.Response, err error) {
	client := resty.New()
	url := fmt.Sprintf("%s/%s", APIUrl, creditCardID)
	body := checkOutRequest{
		Amount:   amount,
		Currency: currency,
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return res, err
	}

	return normalizeCheckoutResp(resp.Body())

}

func normalizeCheckoutResp(resBody []byte) (res response.Response, err error) {
	cRes := checkOutResponse{}
	err = json.Unmarshal(resBody, &cRes)
	if err != nil {
		return res, err
	}
	switch cRes.Status {
	case RespStatusSuccessed:
		return response.NewSuccessfulResponse(cRes), nil
	case RespStatusFailed:
		return response.NewFailureResponse(cRes.Outcome.Reason, cRes), nil
	}

	switch cRes.Outcome.Type {
	case RespOutcomeTypeBlocked:
		return response.NewBlockedResponse(cRes), nil
	case RespOutcomeTypeIssuerDeclined:
		reason := cRes.Outcome.Reason
		switch reason {
		case RespOutcomeDeclinedReasonStolenCard:
			return response.NewStolenCardResponse(cRes), nil
		case RespOutcomeDeclinedReasonInsufficientFunds:
			return response.NewInsufficientFundsResponse(cRes), nil
		case RespOutcomeDeclinedReasonLostCard:
			return response.NewLostCardResponse(cRes), nil
		case RespOutcomeDeclinedReasonHighestRiskLevel:
			return response.NewHighRiskLevelResponse(cRes), nil
		default:
			return response.NewFailureResponse(reason, cRes), nil
		}
	}

	switch cRes.Error.Type {
	case RespErrorTypeInvalidReq:
		if cRes.Error.Message == RespErrorInvalidCreditCardMsg {
			return response.NewInvalidCreditCardResponse(cRes), nil
		}
		return response.NewFailureResponse(cRes.Error.Message, cRes), nil
	}

	return response.NewUnknownResponse(cRes), nil
}
