package errors

var (
	UnknownError = NewBaseError().Kind("unknown error")

	BadRequest          = NewBaseError("bad request")
	NotFound            = NewBaseError("not found")
	InternalServerError = NewBaseError("internal server error")

	ValidationError = NewBaseError("validation error")
)
