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
func Spliit_NewSender(baseUrl *C.char, outHandle *C.uint64_t) C.int {
	if baseUrl == nil {
		setLastError(nullPointerError("base_url"))
		return SpliitError
	}
	if outHandle == nil {
		setLastError(nullPointerError("out_handle"))
		return SpliitError
	}

	sender := spliit.NewHTTPSender(C.GoString(baseUrl), http.DefaultClient)

	*outHandle = C.uint64_t(registerHandle(sender))
	clearLastError()
	return SpliitSuccess
}
