package http

var ResponseCodes = map[string]int{
	"OK":         200,
	"CREATED":    201,
	"NO_CONTENT": 204,
}

type HttpResponse struct {
	StatusCode int
	Payload    interface{}
}

func (r *DefaultResponse) Ok(json interface{}) HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["OK"],
		Payload:    json,
	}
}

func (r *DefaultResponse) Created(json interface{}) HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["CREATED"],
		Payload:    json,
	}
}

func (r *DefaultResponse) NoContent() HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["NO_CONTENT"],
		Payload:    nil,
	}
}
