package registry

import (
	"context"
	"encoding/json"

	"go.chrastecky.dev/spliit-api/spliit"
	"go.chrastecky.dev/spliit-api/spliit/endpoint"
)

var registry = make(map[string]JsonDispatcher)

type JsonDispatcher func(ctx context.Context, raw json.RawMessage) spliit.Call

func registerDispatcher[TIn any, TOut any](endpoint endpoint.Endpoint[TIn, TOut]) {
	name := endpoint.Name()
	dispatcher := func(ctx context.Context, raw json.RawMessage) spliit.Call {
		return spliit.NewCall(endpoint, raw)
	}

	registry[name] = dispatcher
}

func FindByName(name string) (JsonDispatcher, bool) {
	dispatcher, ok := registry[name]
	return dispatcher, ok
}

func init() {
	registerDispatcher(&endpoint.CreateExpense{})
	registerDispatcher(&endpoint.CreateGroup{})
	registerDispatcher(&endpoint.DeleteExpense{})
	registerDispatcher(&endpoint.GetExpense{})
	registerDispatcher(&endpoint.GetGroup{})
	registerDispatcher(&endpoint.GetGroupDetails{})
	registerDispatcher(&endpoint.GetStats{})
	registerDispatcher(&endpoint.ListActivities{})
	registerDispatcher(&endpoint.ListBalances{})
	registerDispatcher(&endpoint.ListCategories{})
	registerDispatcher(&endpoint.ListExpenses{})
	registerDispatcher(&endpoint.ListGroups{})
	registerDispatcher(&endpoint.UpdateGroup{})
}
