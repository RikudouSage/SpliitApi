package main

/*
#include <stdint.h>
*/
import "C"
import "go.chrastecky.dev/spliit-api/spliit"

//export Spliit_NewClient
func Spliit_NewClient() C.uint64_t {
	client := spliit.NewClient()

	return C.uint64_t(registerHandle(client))
}
