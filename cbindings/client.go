package main

/*
#include <stdint.h>
#include <stdlib.h>

#ifndef SPLIIT_CLIENT_CGO_PREAMBLE
#define SPLIIT_CLIENT_CGO_PREAMBLE 1

typedef struct {
    const char* endpoint;
    const char* error;   // NULL if no error
    const char* result;  // JSON string, may be NULL
} SpliitResult;

#endif
*/
import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	"unsafe"

	"go.chrastecky.dev/spliit-api/cbindings/registry"
	"go.chrastecky.dev/spliit-api/spliit"
)

type jsonCallEnvelope struct {
	Endpoint string          `json:"endpoint"`
	Input    json.RawMessage `json:"input"`
}

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

//export Spliit_SendRequests
func Spliit_SendRequests(clientHandle C.uint64_t, jsonCalls *C.char, outResults **C.SpliitResult, outCount *C.size_t) C.int {
	if jsonCalls == nil {
		setLastError(nullPointerError("jsonCalls"))
		return SpliitError
	}
	if outResults == nil {
		setLastError(nullPointerError("outResults"))
		return SpliitError
	}
	if outCount == nil {
		setLastError(nullPointerError("outCount"))
		return SpliitError
	}

	client, err := getHandleObj[spliit.Client](pointerHandle(clientHandle))
	if err != nil {
		setLastError(err)
		return SpliitError
	}

	rawJson := C.GoString(jsonCalls)
	var calls []jsonCallEnvelope
	if err = json.Unmarshal([]byte(rawJson), &calls); err != nil {
		setLastError(err)
		return SpliitError
	}

	if len(calls) == 0 {
		setLastErrorMessage("the calls cannot be empty")
		return SpliitError
	}

	ctx := context.Background()
	goCalls := make([]spliit.Call, 0, len(calls))
	for _, envelope := range calls {
		dispatcher, exists := registry.FindByName(envelope.Endpoint)
		if !exists {
			setLastErrorMessage(fmt.Sprintf("endpoint %s does not exist", envelope.Endpoint))
			return SpliitError
		}
		goCalls = append(goCalls, dispatcher(ctx, envelope.Input))
	}

	_, err = client.SendRequests(ctx, goCalls...)
	if err != nil {
		setLastError(err)
		return SpliitError
	}

	callCount := len(goCalls)
	size := C.size_t(unsafe.Sizeof(C.SpliitResult{}))
	memory := C.calloc(C.size_t(callCount), size)

	if memory == nil {
		setLastErrorMessage("out of memory")
		return SpliitError
	}

	resultSlice := unsafe.Slice((*C.SpliitResult)(memory), callCount)

	freePartialResults := func() {
		for i := 0; i < callCount; i++ {
			if resultSlice[i].endpoint != nil {
				C.free(unsafe.Pointer(resultSlice[i].endpoint))
			}
			if resultSlice[i].error != nil {
				C.free(unsafe.Pointer(resultSlice[i].error))
			}
			if resultSlice[i].result != nil {
				C.free(unsafe.Pointer(resultSlice[i].result))
			}
		}

		C.free(memory)
	}

	for i, call := range goCalls {
		resultSlice[i].endpoint = C.CString(call.EndpointName())
		rawResult, ok := call.(spliit.RawResultCall)
		if !ok {
			freePartialResults()
			setLastErrorMessage("internal error: call does not implement RawResultCall")
			return SpliitError
		}

		rawOutputJson, err := rawResult.RawJson()
		if err != nil {
			freePartialResults()
			setLastError(err)
			return SpliitError
		}

		if rawResult.ErrValue() != nil {
			resultSlice[i].error = C.CString(rawResult.ErrValue().Error())
			resultSlice[i].result = nil
		} else {
			resultSlice[i].result = C.CString(rawOutputJson)
		}
	}

	*outResults = (*C.SpliitResult)(memory)
	*outCount = C.size_t(callCount)

	clearLastError()
	return SpliitSuccess
}

//export Spliit_FreeResults
func Spliit_FreeResults(results *C.SpliitResult, count C.size_t) {
	if results == nil {
		return
	}

	slice := unsafe.Slice(results, int(count))
	for i := range slice {
		if slice[i].endpoint != nil {
			C.free(unsafe.Pointer(slice[i].endpoint))
		}
		if slice[i].error != nil {
			C.free(unsafe.Pointer(slice[i].error))
		}
		if slice[i].result != nil {
			C.free(unsafe.Pointer(slice[i].result))
		}
	}
	C.free(unsafe.Pointer(results))
}
