package main

/*
#include <stdint.h>
*/
import "C"
import (
	"errors"
	"io"
	"sync"
)

type pointerHandle uint64

var (
	handlesMutex sync.Mutex
	handles                    = make(map[pointerHandle]any)
	nextHandleId pointerHandle = 1
)

func registerHandle[TType any](client TType) pointerHandle {
	handlesMutex.Lock()
	defer handlesMutex.Unlock()

	id := nextHandleId
	nextHandleId++
	handles[id] = client
	return id
}

func getHandleObj[TType any](id pointerHandle) (TType, error) {
	handlesMutex.Lock()
	defer handlesMutex.Unlock()

	var zero TType

	obj, ok := handles[id]
	if !ok {
		return zero, errors.New("no value registered for this id")
	}

	typedObj, ok := obj.(TType)
	if !ok {
		return zero, errors.New("handle type mismatch")
	}

	return typedObj, nil
}

func unregisterHandle(id pointerHandle) error {
	handlesMutex.Lock()

	obj, ok := handles[id]
	if !ok {
		handlesMutex.Unlock()
		return errors.New("no object registered for this id")
	}

	delete(handles, id)
	handlesMutex.Unlock()

	if closer, ok := obj.(io.Closer); ok {
		closer.Close()
	}

	return nil
}

//export Spliit_CloseHandle
func Spliit_CloseHandle(handle C.uint64_t) C.int {
	err := unregisterHandle(pointerHandle(handle))
	if err != nil {
		setLastError(err)
		return SpliitError
	}

	clearLastError()
	return SpliitSuccess
}
