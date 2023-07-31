package client

import (
	"encoding/json"
	"fmt"
)

type RequestBuilder interface {
	GetRequest(*Client) (*Request, error)
}

type Request struct {
	body *[]byte
}

func NewRequest(body *[]byte) *Request {
	return &Request{body}
}

func (at *Request) Unmarshal(dest interface{}) error {
	err := json.Unmarshal(*at.body, dest)
	if err != nil {
		return fmt.Errorf("error unmarshalling request body: %w", err)
	}

	return nil
}
