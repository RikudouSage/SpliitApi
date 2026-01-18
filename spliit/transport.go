package spliit

import (
	"encoding/json"
	"fmt"
)

type OutboundRequest struct {
	Endpoint string          `json:"endpoint"`
	Input    json.RawMessage `json:"input"`
	Mutates  bool            `json:"-"`
}

type InboundResponse struct {
	Endpoint string          `json:"endpoint"`
	Result   json.RawMessage `json:"result,omitempty"`
	Error    *RemoteError    `json:"error,omitempty"`
}

type RemoteError struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (err *RemoteError) Error() string {
	if err == nil {
		return "<nil>"
	}
	if err.Code == "" {
		return err.Message
	}
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}
