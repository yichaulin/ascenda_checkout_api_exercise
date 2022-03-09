package response

type Response struct {
	Status        string      `json:"status"`
	FailureReason *string     `json:"failureReason"`
	Outcome       interface{} `json:"outcome"`
}

const (
	StatusSuccessful = "Successful"
	StatusFailure    = "Failure"
	StatusUnknown    = "Unknown"

	FailureReasonStolenCard        = "Stolen card"
	FailureReasonInsufficientFunds = "Insufficient funds"
	FailureReasonLostCard          = "Lost card"
	FailureReasonInvalidCreditCard = "Invalid credit card number"
	FailureReasonBlocked           = "Fradulent. Blocked"
	FailureReasonHighRiskLevel     = "High Risk Level"
)

func NewSuccessfulResponse(outcome interface{}) Response {
	return Response{
		Status:  StatusSuccessful,
		Outcome: outcome,
	}
}

func NewFailureResponse(reason string, outcome interface{}) Response {
	return Response{
		Status:        StatusFailure,
		FailureReason: &reason,
		Outcome:       outcome,
	}
}

func NewStolenCardResponse(outcome interface{}) Response {
	return NewFailureResponse(FailureReasonStolenCard, outcome)
}

func NewInsufficientFundsResponse(outcome interface{}) Response {
	return NewFailureResponse(FailureReasonInsufficientFunds, outcome)
}

func NewLostCardResponse(outcome interface{}) Response {
	return NewFailureResponse(FailureReasonLostCard, outcome)
}

func NewInvalidCreditCardResponse(outcome interface{}) Response {
	return NewFailureResponse(FailureReasonInvalidCreditCard, outcome)
}

func NewBlockedResponse(outcome interface{}) Response {
	return NewFailureResponse(FailureReasonBlocked, outcome)
}

func NewHighRiskLevelResponse(outcome interface{}) Response {
	return NewFailureResponse(FailureReasonHighRiskLevel, outcome)
}

func NewUnknownResponse(outcome interface{}) Response {
	return Response{
		Status:  StatusUnknown,
		Outcome: outcome,
	}
}
