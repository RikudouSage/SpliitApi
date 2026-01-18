package spliit

import (
	"encoding/json"
	"go.chrastecky.dev/spliit-api/spliit/endpoint"
	"go.chrastecky.dev/spliit-api/spliit/shape"
	"reflect"
)

// Request ties an endpoint to a concrete (possibly untyped) input.
type Request[TInput any, TOutput any] struct {
	Endpoint endpoint.Endpoint[TInput, TOutput]
	Input    any
}

// NewRequest creates a typed request for an endpoint.
func NewRequest[TInput any, TOutput any](ep endpoint.Endpoint[TInput, TOutput], input any) Request[TInput, TOutput] {
	return Request[TInput, TOutput]{
		Endpoint: ep,
		Input:    input,
	}
}

// ValidateInput validates and converts the input to the endpoint's input type.
func (req Request[TInput, TOutput]) ValidateInput() (TInput, error) {
	decoded, err := decodeStrict[TInput](req.Input)
	if err != nil {
		return decoded, err
	}
	applyDefaults(&decoded)
	return decoded, nil
}

// DecodeOutput validates and converts the output to the endpoint's output type.
func (req Request[TInput, TOutput]) DecodeOutput(data json.RawMessage) (TOutput, error) {
	return decodeStrict[TOutput](data)
}

func applyDefaults[T any](value *T) {
	if value == nil {
		return
	}

	if def, ok := any(value).(shape.RequestDefaults); ok {
		def.ApplyDefaults()
		return
	}

	if isNilValue(*value) {
		return
	}

	if def, ok := any(*value).(shape.RequestDefaults); ok {
		def.ApplyDefaults()
	}
}

func isNilValue[T any](value T) bool {
	valueType := reflect.TypeOf(value)
	if valueType == nil {
		return true
	}
	//nolint:exhaustive // only nil-able kinds should return true here
	switch valueType.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return reflect.ValueOf(value).IsNil()
	default:
		return false
	}
}
