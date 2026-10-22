package spliit

import (
	"context"
	"fmt"
	"net/http"

	"go.chrastecky.dev/spliit-api/cfg"
)

// Client sends batches of Spliit calls.
type Client interface {
	// SendRequests sends one or more calls in a single batch request.
	// The returned slice contains the same call objects passed in calls, updated in place with results and per-call errors.
	SendRequests(ctx context.Context, calls ...Call) ([]Call, error)
}

type client struct {
	sender Sender
}

// NewClient constructs a client that uses the default HTTPS sender.
func NewClient() Client {
	return NewClientWithSender(nil)
}

// NewClientWithSender constructs a client that uses sender.
// If sender is nil, a default HTTPS sender is created.
func NewClientWithSender(sender Sender) Client {
	if sender == nil {
		sender = NewHTTPSender(cfg.DefaultApiURL, http.DefaultClient)
	}

	return &client{sender: sender}
}

// SendRequests sends calls and attaches each response to the matching call.
// The returned slice contains the same call objects passed in calls, updated in place with results and per-call errors.
func (receiver *client) SendRequests(ctx context.Context, calls ...Call) ([]Call, error) {
	if len(calls) == 0 {
		return nil, nil
	}

	outbound := make([]OutboundRequest, 0, len(calls))
	for _, call := range calls {
		payload, err := call.encodeInput()
		if err != nil {
			return nil, err
		}
		outbound = append(outbound, OutboundRequest{
			Endpoint: call.EndpointName(),
			Input:    payload,
			Mutates:  call.mutates(),
		})
	}

	inbound := make([]InboundResponse, len(calls))
	for _, mutates := range []bool{false, true} {
		indices := make([]int, 0, len(outbound))
		batch := make([]OutboundRequest, 0, len(outbound))
		for i, request := range outbound {
			if request.Mutates == mutates {
				indices = append(indices, i)
				batch = append(batch, request)
			}
		}

		if len(batch) == 0 {
			continue
		}

		responses, err := receiver.sender.SendBatch(ctx, batch)
		if err != nil {
			return nil, err
		}

		if len(responses) != len(batch) {
			return nil, fmt.Errorf("response count mismatch: got %d, want %d", len(responses), len(batch))
		}

		for i, response := range responses {
			inbound[indices[i]] = response
		}
	}

	for i, resp := range inbound {
		call := calls[i]
		if resp.Endpoint != call.EndpointName() {
			return nil, fmt.Errorf("response[%d] endpoint mismatch: got %q, want %q", i, resp.Endpoint, call.EndpointName())
		}
		if err := call.applyResponse(resp); err != nil {
			call.setError(err)
		}
	}

	return calls, nil
}
