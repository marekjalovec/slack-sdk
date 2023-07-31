package payloads

import (
	"encoding/json"
	"fmt"

	"github.com/marekjalovec/slack-sdk/views"
)

type ResponseActionPayload struct {
	ResponseAction string      `json:"response_action"`
	View           *views.View `json:"view"`
}

func NewResponseActionPayload(view *views.View) *ResponseActionPayload {
	return &ResponseActionPayload{
		ResponseAction: "update",
		View:           view,
	}
}

func (at *ResponseActionPayload) Stringify() (string, error) {
	s, err := json.Marshal(at)
	if err != nil {
		return "", fmt.Errorf("error marshalling response action payload: %w", err)
	}

	return string(s), nil
}
