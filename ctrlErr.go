package errors

type ControllerError struct {
	BaseError
	jsonResError map[string]interface{}
	httpCode     int
}

// New Controller Error
func NewCtlError(be *BaseError) *ControllerError {
	return &ControllerError{
		BaseError: *be,
	}
}

// // NEED BETTER NAMES
// func NewControllerErrorWithRoot(kind, rootCause string) *ControllerError {
// 	return &ControllerError{
// 		BaseError: *BaseWithRoot(kind, rootCause),
// 	}
// }

func (ce *ControllerError) JsonRes() map[string]interface{} {

	switch ce.BaseError.kind {
	//# I want to change this
	case "not found":
		ce.httpCode = 404
	case "bad request":
		ce.httpCode = 400
	default:
		ce.httpCode = 500
	}

	ce.jsonResError = map[string]interface{}{"error": ce.Error()}
	return ce.jsonResError
}

func (ce *ControllerError) GetHttpCode() int {
	return ce.httpCode
}
