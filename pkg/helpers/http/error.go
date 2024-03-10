package http

var errorCodes = map[string]int{
	"BAD_REQUEST":           400,
	"UNAUTHORIZED":          401,
	"FORBIDDEN":             403,
	"NOT_FOUND":             404,
	"INTERNAL_SERVER_ERROR": 500,
	"DUPLICATED":            409,
	"UNPROCESSABLE_ENTITY":  422,
	"TOO_MANY_REQUESTS":     429,
	"SERVICE_UNAVAILABLE":   503,
	"NOT_IMPLEMENTED":       501,
}

type HttpError struct {
	StatusCode int
	Message    string
	error
}

func (r *DefaultError) BadRequest(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["BAD_REQUEST"],
		Message:    message,
	}
}

func (r *DefaultError) Unauthorized(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["UNAUTHORIZED"],
		Message:    message,
	}
}

func (r *DefaultError) Forbidden(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["FORBIDDEN"],
		Message:    message,
	}
}

func (r *DefaultError) NotFound(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["NOT_FOUND"],
		Message:    message,
	}
}

func (r *DefaultError) InternalServerError(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["INTERNAL_SERVER_ERROR"],
		Message:    message,
	}
}

func (r *DefaultError) Duplicated(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["DUPLICATED"],
		Message:    message,
	}
}

func (r *DefaultError) UnprocessableEntity(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["UNPROCESSABLE_ENTITY"],
		Message:    message,
	}
}

func (r *DefaultError) TooManyRequests(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["TOO_MANY_REQUESTS"],
		Message:    message,
	}
}

func (r *DefaultError) ServiceUnavailable(message string) HttpError {
	return HttpError{
		StatusCode: errorCodes["SERVICE_UNAVAILABLE"],
		Message:    message,
	}
}
