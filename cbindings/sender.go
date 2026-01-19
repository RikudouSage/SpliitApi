package main

/*
#include <stdint.h>
*/
import "C"
import (
	"net/http"

	"go.chrastecky.dev/spliit-api/spliit"
)

//export Spliit_NewSender
func Spliit_NewSender(baseUrl *C.char) C.uint64_t {
	sender := spliit.NewHTTPSender(C.GoString(baseUrl), http.DefaultClient)

	return C.uint64_t(registerHandle(sender))
}
