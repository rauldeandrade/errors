package main

type erratum interface {
	BaseError | ControllerError
}

func NewErr[T erratum]() *T {
	t := new(T)
	return t
}
