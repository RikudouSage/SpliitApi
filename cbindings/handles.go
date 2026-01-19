package main

import (
	"errors"
	"sync"

	"go.chrastecky.dev/spliit-api/spliit"
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

func getClient(id pointerHandle) (spliit.Client, error) {
	handlesMutex.Lock()
	defer handlesMutex.Unlock()

	obj, exists := handles[id]
	if !exists {
		return nil, errors.New("no client registered for this id")
	}

	if client, ok := obj.(spliit.Client); ok {
		return client, nil
	}

	return nil, errors.New("no client registered for this id")
}

func unregisterClient(id pointerHandle) {
	handlesMutex.Lock()
	defer handlesMutex.Unlock()
	delete(handles, id)
}
