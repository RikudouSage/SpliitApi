package main

/*
#include <stdint.h>
*/
import "C"
import (
	"sync"

	"go.chrastecky.dev/spliit-api/spliit"
)

type clientHandle uint64

var (
	clientsMutex sync.Mutex
	clients                   = make(map[clientHandle]spliit.Client)
	nextClientID clientHandle = 1
)

func registerClient(client spliit.Client) clientHandle {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	id := nextClientID
	nextClientID++
	clients[id] = client
	return id
}

func getClient(id clientHandle) spliit.Client {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	return clients[id]
}

func unregisterClient(id clientHandle) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	delete(clients, id)
}

//export Spliit_NewClient
func Spliit_NewClient() C.uint64_t {
	client := spliit.NewClient()
	handle := registerClient(client)

	return C.uint64_t(handle)
}

//export Spliit_CloseClient
func Spliit_CloseClient(handle C.uint64_t) {
	unregisterClient(clientHandle(handle))
}
