package errors

import "fmt"

type BaseError struct {
	kind      string // type of error (ex. not found)
	rootCause string // original reason for the error (ex. no document in db)
	message   string // additional information
	errChain  error  // chain of upstream errors that caused this error
}

//func NewBaseError(kind, rootCause, message)
//
//New generic error, which can serve as a base for other errors.
//This error contains (kind, root cause, message) string, (error chain) error
func NewBaseError(fields ...string) *BaseError {
	//I don't like (fields ...string)

	b := &BaseError{}

	switch len(fields) {
	case 1:
		b.kind = fields[0]
	case 2:
		b.rootCause = fields[1]
	case 3:
		b.message = fields[2]
	default:
		return b
	}
	return b
}

func (e *BaseError) Error() string {
	if e.errChain != nil {
		return e.errChain.Error()
	}

	str := ""
	if e.kind != "" {
		str += ": " + e.kind
	}
	if e.message != "" {
		str += ": " + e.message
	}
	if e.rootCause != "" {
		if str != "" {
			str += ": root cause: " + e.rootCause
		} else {
			str += "root cause: " + e.rootCause
		}
	}
	return fmt.Sprintf("ERROR %s", str)
}

func (e *BaseError) Kind(kind ...string) string {
	if kind != nil {
		e.kind = kind[0]
	}
	return e.kind
}

func (e *BaseError) RootCause(rootCause ...string) string {
	if rootCause != nil {
		e.rootCause = rootCause[0]
	}
	return e.rootCause
}

func (e *BaseError) Message(msg ...string) string {
	if msg != nil {
		e.message = msg[0]
	}
	return e.message
}

func (e *BaseError) Wrap(err error) error {
	e.errChain = fmt.Errorf("%s :: %w", e.Error(), err)
	return e
}
