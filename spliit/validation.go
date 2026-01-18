package spliit

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"reflect"
)

func decodeStrict[T any](input any) (T, error) {
	var zero T
	if input == nil {
		if allowsNil(zero) {
			return zero, nil
		}
		return zero, errors.New("input is nil")
	}
	if value, ok := input.(T); ok {
		return value, nil
	}

	data, err := toJSONBytes(input)
	if err != nil {
		return zero, err
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&zero); err != nil {
		return zero, err
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return zero, errors.New("unexpected additional JSON values")
	}

	return zero, nil
}

func toJSONBytes(input any) ([]byte, error) {
	switch value := input.(type) {
	case json.RawMessage:
		return value, nil
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	default:
		return json.Marshal(value)
	}
}

func allowsNil[T any](value T) bool {
	valueType := reflect.TypeOf(value)
	if valueType == nil {
		return true
	}
	switch valueType.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return true
	default:
		return false
	}
}
