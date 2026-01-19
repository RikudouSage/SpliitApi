package spliit

import (
	"context"
	"fmt"
	"net/http"

	"go.chrastecky.dev/spliit-api/cfg"
)

type Client interface {
	SendRequests(ctx context.Context, calls ...Call) ([]Call, error)
}

type client struct {
	sender Sender
}

func NewClient() Client {
	return NewClientWithSender(nil)
}

func NewClientWithSender(sender Sender) Client {
	if sender == nil {
		sender = NewHTTPSender(cfg.DefaultApiUrl, http.DefaultClient)
	}

	return &client{sender: sender}
}

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

	inbound, err := receiver.sender.SendBatch(ctx, outbound)
	if err != nil {
		return nil, err
	}
	if len(inbound) != len(calls) {
		return nil, fmt.Errorf("response count mismatch: got %d, want %d", len(inbound), len(calls))
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
