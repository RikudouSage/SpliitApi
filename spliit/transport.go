package spliit

import (
	"encoding/json"
	"fmt"
)

// OutboundRequest is the transport representation of a single outgoing call.
type OutboundRequest struct {
	Endpoint string          `json:"endpoint"`
	Input    json.RawMessage `json:"input"`
	Mutates  bool            `json:"-"`
}

// InboundResponse is the transport representation of a single response item.
type InboundResponse struct {
	Endpoint string          `json:"endpoint"`
	Result   json.RawMessage `json:"result,omitempty"`
	Error    *RemoteError    `json:"error,omitempty"`
}

// RemoteError describes an error returned by the remote tRPC API.
type RemoteError struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

// Error formats the remote error as an error string.
func (err *RemoteError) Error() string {
	if err == nil {
		return "<nil>"
	}
	if err.Code == "" {
		return err.Message
	}
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}
