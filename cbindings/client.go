package main

/*
#include <stdint.h>
*/
import "C"
import "go.chrastecky.dev/spliit-api/spliit"

//export Spliit_NewClient
func Spliit_NewClient() C.uint64_t {
	client := spliit.NewClient()
	handle := registerHandle(client)

	return C.uint64_t(handle)
}

//export Spliit_CloseClient
func Spliit_CloseClient(handle C.uint64_t) {
	unregisterClient(pointerHandle(handle))
}
