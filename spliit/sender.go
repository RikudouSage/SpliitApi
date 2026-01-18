package spliit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Sender interface {
	SendBatch(ctx context.Context, requests []OutboundRequest) ([]InboundResponse, error)
}

// HTTPSender calls a tRPC HTTP batch endpoint.
type HTTPSender struct {
	baseURL string
	client  *http.Client
}

// NewHTTPSender creates an HTTP sender for a tRPC API base URL (e.g. https://spliit.app/api/trpc).
// If httpClient is nil, http.DefaultClient is used.
func NewHTTPSender(baseURL string, httpClient *http.Client) *HTTPSender {
	return &HTTPSender{
		baseURL: strings.TrimRight(baseURL, "/"),
		client:  httpClient,
	}
}

func (sender *HTTPSender) SendBatch(ctx context.Context, requests []OutboundRequest) ([]InboundResponse, error) {
	if len(requests) == 0 {
		return nil, nil
	}
	if sender.baseURL == "" {
		return nil, errors.New("http sender baseURL is empty")
	}
	httpClient := sender.client
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	endpointNames := make([]string, 0, len(requests))
	input := make(map[string]trpcInputEnvelope, len(requests))
	for i, req := range requests {
		endpointNames = append(endpointNames, req.Endpoint)
		input[fmt.Sprintf("%d", i)] = trpcInputEnvelope{JSON: req.Input}
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to encode input payload: %w", err)
	}

	usePost := anyMutates(requests)
	batchURL, err := buildBatchURL(sender.baseURL, endpointNames, inputBytes, usePost)
	if err != nil {
		return nil, err
	}

	method := http.MethodGet
	var reqBody io.Reader
	if usePost {
		method = http.MethodPost
		reqBody = strings.NewReader(string(inputBytes))
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, batchURL, reqBody)
	if err != nil {
		return nil, err
	}
	if usePost {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("http status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return decodeBatchResponse(body, requests)
}

func buildBatchURL(baseURL string, endpoints []string, inputBytes []byte, usePost bool) (string, error) {
	if len(endpoints) == 0 {
		return "", errors.New("no endpoints for batch")
	}

	parsed, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	path := strings.TrimRight(parsed.Path, "/") + "/" + strings.Join(endpoints, ",")
	parsed.Path = path

	query := parsed.Query()
	query.Set("batch", "1")
	if !usePost {
		query.Set("input", string(inputBytes))
	}
	parsed.RawQuery = query.Encode()

	return parsed.String(), nil
}

func anyMutates(requests []OutboundRequest) bool {
	for _, req := range requests {
		if req.Mutates {
			return true
		}
	}
	return false
}

type trpcInputEnvelope struct {
	JSON json.RawMessage `json:"json"`
}

func decodeBatchResponse(body []byte, requests []OutboundRequest) ([]InboundResponse, error) {
	var envelopes []trpcResponseEnvelope
	if err := json.Unmarshal(body, &envelopes); err != nil {
		return nil, fmt.Errorf("failed to decode batch response: %w", err)
	}

	if len(envelopes) != len(requests) {
		return nil, fmt.Errorf("response count mismatch: got %d, want %d", len(envelopes), len(requests))
	}

	responses := make([]InboundResponse, 0, len(envelopes))
	for i, env := range envelopes {
		resp := InboundResponse{Endpoint: requests[i].Endpoint}
		if env.Error != nil {
			resp.Error = &RemoteError{
				Code:    formatErrorCode(env.Error.Code),
				Message: env.Error.Message,
				Data:    env.Error.Data,
			}
			responses = append(responses, resp)
			continue
		}

		if env.Result != nil && env.Result.Data != nil && len(env.Result.Data.JSON) > 0 {
			resp.Result = env.Result.Data.JSON
		} else {
			resp.Result = json.RawMessage("null")
		}

		responses = append(responses, resp)
	}

	return responses, nil
}

func formatErrorCode(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var codeString string
	if err := json.Unmarshal(raw, &codeString); err == nil {
		return codeString
	}
	var codeNumber float64
	if err := json.Unmarshal(raw, &codeNumber); err == nil {
		return fmt.Sprintf("%g", codeNumber)
	}
	return string(raw)
}

type trpcResponseEnvelope struct {
	Result *trpcResponseResult `json:"result,omitempty"`
	Error  *trpcResponseError  `json:"error,omitempty"`
}

type trpcResponseResult struct {
	Data *trpcResponseData `json:"data,omitempty"`
}

type trpcResponseData struct {
	JSON json.RawMessage `json:"json"`
}

type trpcResponseError struct {
	Message string          `json:"message"`
	Code    json.RawMessage `json:"code"`
	Data    json.RawMessage `json:"data,omitempty"`
}
