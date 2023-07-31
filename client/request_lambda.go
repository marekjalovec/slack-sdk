package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type LambdaRequest struct {
	RawBody         string               `json:"body"`
	Headers         LambdaRequestHeaders `json:"headers"`
	IsBase64Encoded bool                 `json:"isBase64Encoded"`
}

type LambdaRequestHeaders struct {
	Signature        string `json:"x-slack-signature"`
	RequestTimestamp string `json:"x-slack-request-timestamp"`
}

func (at *LambdaRequest) GetRequest(client *Client) (*Request, error) {
	if at.IsBase64Encoded {
		var rawBody, err = base64.StdEncoding.DecodeString(at.RawBody)
		if err != nil {
			return nil, fmt.Errorf("request body decoding failed: %w", err)
		}

		at.RawBody = string(rawBody)
		at.IsBase64Encoded = false
	}

	err := client.validateEvent(at.RawBody, at.Headers.RequestTimestamp, at.Headers.Signature)
	if err != nil {
		return nil, err
	}

	body, err := normalizeRequestBody(at.RawBody)
	if err != nil {
		return nil, err
	}

	return NewRequest(body), nil
}

func normalizeRequestBody(s string) (*[]byte, error) {
	s = strings.TrimSpace(s)

	if strings.HasPrefix(s, "payload=") {
		params, err := url.ParseQuery(s)
		if err != nil {
			return nil, fmt.Errorf("error normalizing request body from payload query: %w", err)
		}
		payload, ok := params["payload"]
		if ok && len(payload) == 1 {
			body := []byte(payload[0])
			return &body, nil
		} else {
			return nil, fmt.Errorf("error normalizing request body from payload query: %w", err)
		}
	} else if isJSON(s) {
		body := []byte(s)
		return &body, nil
	} else if isQuery(s) {
		query, _ := url.ParseQuery(s)
		val := map[string]string{}
		for k, v := range query {
			val[k] = v[0]
		}
		body, err := json.Marshal(val)
		if err != nil {
			return nil, fmt.Errorf("error normalizing request body from query: %w", err)
		}
		return &body, nil
	}

	return nil, fmt.Errorf("unknown request body format")
}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func isQuery(str string) bool {
	_, err := url.ParseQuery(str)
	return err == nil
}
