package client

import (
	"encoding/json"
	"fmt"
)

type RequestBuilder interface {
	GetRequest(*Client) (*Request, error)
}

type Request struct {
	Type    RequestType
	rawBody *string
	Body    *[]byte
}

type RequestType string

const (
	RequestTypeEvent       RequestType = "event"
	RequestTypeInteraction RequestType = "interaction"
	RequestTypeCommand     RequestType = "command"
)

func (at *Request) Unmarshal(dest interface{}) error {
	err := json.Unmarshal(*at.Body, dest)
	if err != nil {
		return fmt.Errorf("error unmarshalling request body: %w", err)
	}

	return nil
}
