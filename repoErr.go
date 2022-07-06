package errors

type RepoError struct {
	BaseError
}

// New Service Error
func NewRepoError(be *BaseError) *RepoError {
	return &RepoError{
		BaseError: *be,
	}
}

// // NEED BETTER NAMES
// func NewServiceErrorWithRoot(kind, rootCause string) *ServiceError {
// 	return &ServiceError{
// 		BaseError: *BaseWithRoot(kind, rootCause),
// 	}
// }

// func (ce *ServiceError) JsonRes() map[string]interface{} {

// 	ce.jsonResError = map[string]interface{}{"error": ce.Error()}
// 	return ce.jsonResError
// }

// func (ce *ServiceError) GetHttpCode() int {
// 	switch ce.BaseError.kind {
// 	//# I want to change this
// 	case "not found":
// 		ce.httpCode = 404
// 	case "bad request":
// 		ce.httpCode = 400
// 	default:
// 		ce.httpCode = 500
// 	}
// 	return ce.httpCode
// }
