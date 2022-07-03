package errors

import "fmt"

type BaseError struct {
	kind      string // type of error (ex. not found)
	rootCause string // original reason for the error (ex. no document in db)
	message   string // additional information
	errChain  error  // chain of upstream errors that caused this error
}

func New() *BaseError {
	return &BaseError{}
}

func (e *BaseError) Error() string {
	if e.errChain != nil {
		return e.errChain.Error()
	}

	str := ""
	if e.kind != "" {
		str += e.kind
	}
	if e.message != "" {
		str += ": " + e.message
	}
	if e.rootCause != "" {
		if str != "" {
			str += ": root cause:\n\t\t" + e.rootCause
		} else {
			str += "root cause:\n\t\t" + e.rootCause
		}
	}
	return fmt.Sprintf("ERROR:  %s", str)
}

func (e *BaseError) Kind(kind string) string {
	if kind != "" {
		e.kind = kind
	}
	return e.kind
}

func (e *BaseError) RootCause(rootCause string) string {
	if rootCause != "" {
		e.rootCause = rootCause
	}
	return e.rootCause
}

func (e *BaseError) Message(msg string) string {
	if msg != "" {
		e.message = msg
	}
	return e.message
}

func (e *BaseError) Wrap(err error) error {
	e.errChain = fmt.Errorf("%s :: %w", e.Error(), err)
	return e
}
