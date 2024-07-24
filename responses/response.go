package responses

type Fail struct {
	Status bool        `json:"status" extensions:"x-order=0"`
	Code   int         `json:"code" extensions:"x-order=1"`
	Err    ErrorFormat `json:"error" extensions:"x-order=2"`
}

type ErrorFormat struct {
	Message string      `json:"message" extensions:"x-order=0"`
	Field   interface{} `json:"field" extensions:"x-order=1"`
}

func ValidateResponse(msg map[string]string) Fail {
	var request Fail
	request.Code = 422
	request.Err.Message = "validate error"
	request.Err.Field = msg
	request.Status = false
	return request
}