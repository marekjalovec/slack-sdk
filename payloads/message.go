package payloads

import (
	"encoding/json"
	"fmt"

	"github.com/marekjalovec/slack-sdk/blocks"
	"github.com/marekjalovec/slack-sdk/client"
)

type MessagePayload struct {
	ChannelId string                `json:"channel"`
	Timestamp string                `json:"ts,omitempty"`
	Blocks    *blocks.BlocksWrapper `json:"blocks"`
}

type MessageResponse struct {
	Ok        bool   `json:"ok"`
	ChannelId string `json:"channel"`
	Ts        string `json:"ts"`
	Error     string `json:"error"`
}

func NewMessagePayload(channelId string, b *blocks.BlocksWrapper) *MessagePayload {
	if b == nil {
		b = blocks.NewBlocksWrapper()
	}

	return &MessagePayload{
		ChannelId: channelId,
		Blocks:    b,
	}
}

func (at *MessagePayload) AddBlock(block blocks.Block) {
	at.Blocks.AddBlock(block)
}

func (at *MessagePayload) Publish(client *client.Client) (*MessageResponse, error) {
	r, err := client.Post("chat.postMessage", at)
	if err != nil {
		return nil, err
	}

	mr := MessageResponse{}
	err = json.Unmarshal(r, &mr)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling message response: %w", err)
	}

	return &mr, nil
}

func (at *MessagePayload) Update(client *client.Client) (*MessageResponse, error) {
	r, err := client.Post("chat.update", at)
	if err != nil {
		return nil, err
	}

	mr := MessageResponse{}
	err = json.Unmarshal(r, &mr)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling message response: %w", err)
	}

	return &mr, nil
}
