package spliit

import (
	"encoding/json"
	"fmt"

	"go.chrastecky.dev/spliit-api/spliit/endpoint"
)

// Call is the interface used to send mixed-typed requests in a batch.
type Call interface {
	EndpointName() string
	mutates() bool
	encodeInput() (json.RawMessage, error)
	applyResponse(resp InboundResponse) error
	setError(err error)
}

// OutputCall exposes only the output type for easier type assertions.
type OutputCall[TOutput any] interface {
	Output() TOutput
	ErrValue() error
}

type RawResultCall interface {
	RawJson() (string, error)
	ErrValue() error
}

// TypedCall holds the typed response for a request.
type TypedCall[TInput any, TOutput any] struct {
	Request Request[TInput, TOutput]
	Result  TOutput
	Err     error
}

// NewCall creates a TypedCall from an endpoint and input.
func NewCall[TInput any, TOutput any](ep endpoint.Endpoint[TInput, TOutput], input any) *TypedCall[TInput, TOutput] {
	return NewCallWithRequest(NewRequest(ep, input))
}

// NewCallWithRequest creates a TypedCall that can be passed to SendRequests.
func NewCallWithRequest[TInput any, TOutput any](req Request[TInput, TOutput]) *TypedCall[TInput, TOutput] {
	return &TypedCall[TInput, TOutput]{Request: req}
}

func (call *TypedCall[TInput, TOutput]) EndpointName() string {
	return call.Request.Endpoint.Name()
}

func (call *TypedCall[TInput, TOutput]) mutates() bool {
	return call.Request.Endpoint.Mutates()
}

func (call *TypedCall[TInput, TOutput]) Output() TOutput {
	return call.Result
}

func (call *TypedCall[TInput, TOutput]) ErrValue() error {
	return call.Err
}

func (call *TypedCall[TInput, TOutput]) encodeInput() (json.RawMessage, error) {
	validated, err := call.Request.ValidateInput()
	if err != nil {
		return nil, fmt.Errorf("input validation failed for %q: %w", call.EndpointName(), err)
	}

	payload, err := json.Marshal(validated)
	if err != nil {
		return nil, fmt.Errorf("input marshal failed for %q: %w", call.EndpointName(), err)
	}

	return payload, nil
}

func (call *TypedCall[TInput, TOutput]) applyResponse(resp InboundResponse) error {
	if resp.Error != nil {
		return resp.Error
	}

	result, err := call.Request.DecodeOutput(resp.Result)
	if err != nil {
		return fmt.Errorf("output validation failed for %q: %w", call.EndpointName(), err)
	}

	call.Result = result
	return nil
}

func (call *TypedCall[TInput, TOutput]) RawJson() (string, error) {
	raw, err := json.Marshal(call.Result)
	if err != nil {
		return "", fmt.Errorf("output marshal failed for %q: %w", call.EndpointName(), err)
	}

	return string(raw), nil
}

func (call *TypedCall[TInput, TOutput]) setError(err error) {
	call.Err = err
}
