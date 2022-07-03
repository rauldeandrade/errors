package main

import (
	"github.com/rauldeandrade/errors/errors"
)

var (
	UnknownError = errors.NewBaseError().Kind("unknown error")

	BadRequest          = errors.NewBaseError("bad request")
	NotFound            = errors.NewBaseError("not found")
	InternalServerError = errors.NewBaseError("internal server error")
)
