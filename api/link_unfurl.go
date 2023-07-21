package api

import (
	"fmt"

	"github.com/marekjalovec/slack-sdk/blocks"
	"github.com/marekjalovec/slack-sdk/client"
)

func ChatUnfurl(c *client.Client, channel string, messageTs string, unfurls map[string]*blocks.BlocksWrapper) ([]byte, error) {
	payload := map[string]interface{}{
		"channel": channel,
		"ts":      messageTs,
		"unfurls": map[string]interface{}{},
	}

	for url, wrapper := range unfurls {
		payload["unfurls"].(map[string]interface{})[url] = map[string]interface{}{
			"blocks": wrapper,
		}
	}

	post, err := c.Post("chat.unfurl", payload)
	if err != nil {
		return nil, fmt.Errorf("error posting message: %w", err)
	}

	return post, nil
}
