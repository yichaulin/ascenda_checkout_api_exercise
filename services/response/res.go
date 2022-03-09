package response

type Response struct {
	Status        string      `json:"status"`
	FailureReason string      `json:"failureReason"`
	Outcome       interface{} `json:"outcome"`
}

const (
	StatusSuccessed = "Successed"
	StatusFailure   = "Failure"
	StatusUnknown   = "Unknown"
)

func NewFailureResponse(reason string, outcome interface{}) Response {
	return Response{
		Status:        StatusFailure,
		FailureReason: reason,
		Outcome:       outcome,
	}
}

func NewSuccessfulResponse(outcome interface{}) Response {
	return Response{
		Status:  StatusSuccessed,
		Outcome: outcome,
	}
}

func NewUnknownResponse(outcome interface{}) Response {
	return Response{
		Status:  StatusUnknown,
		Outcome: outcome,
	}
}
