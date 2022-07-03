package main

import (
	"fmt"

	"github.com/rauldeandrade/errors/errors"
)

func NewPrint() {
	fmt.Println("hello.")
}

var (
	UnknownError = errors.New().Kind("unknown error")

	BadRequest          = errors.New().Kind("bad request")
	NotFound            = errors.New().Kind("not found")
	InternalServerError = errors.New().Kind("internal server error")
)

//? return layerError(NotFound, "user") "user: not found" "not found: user"
//? return layerError(BadRequest, "incorrect credentials")

//~ KIND DOES NOT GIVE ME ANY UNIQUE USES LIKE LAYERS
// Layers give special methods, but kind doesn't, so why do I need it?
// Kind in these examples is really an HTTP CODE, so it's only useful
// in controller layer

//? Other examples
//? Service layer id validation -> bad request "invalid id"
//?				"bad request: invalid id" | "invalid id: bad request"
//? Repo layer data deleted -> can't update "data deleted"
//?				"can't update: data deleted" | "data deleted: can't update"
//? json.NewEncoder(res).Encode(errors.New("unable to unmarshal comment"))
//? json.NewEncoder(res).Encode(ctrlErr(Unmarshal, "comment"))
//?				"unmarshal error: comment"
