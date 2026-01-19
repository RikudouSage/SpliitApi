package main

/*
#include <stdint.h>
*/
import "C"
import (
	"go.chrastecky.dev/spliit-api/spliit"
)

//export Spliit_NewClient
func Spliit_NewClient(outHandle *C.uint64_t) C.int {
	if outHandle == nil {
		setLastError(nullPointerError("outHandle"))
		return SpliitError
	}

	client := spliit.NewClient()

	*outHandle = C.uint64_t(registerHandle(client))
	clearLastError()
	return SpliitSuccess
}
