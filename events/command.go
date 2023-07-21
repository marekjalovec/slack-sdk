package events

import (
	"strings"

	"github.com/marekjalovec/slack-sdk/client"
)

type Command struct {
	Token               string `json:"token"`
	TeamId              string `json:"team_id"`
	TeamDomain          string `json:"team_domain"`
	ChannelId           string `json:"channel_id"`
	ChannelName         string `json:"channel_name"`
	UserId              string `json:"user_id"`
	UserName            string `json:"user_name"`
	Command             string `json:"command"`
	Text                string `json:"text"`
	ApiAppId            string `json:"api_app_id"`
	IsEnterpriseInstall string `json:"is_enterprise_install"`
	ResponseUrl         string `json:"response_url"`
	TriggerId           string `json:"trigger_id"`
}

func GetCommand(r *client.Request) (*Command, error) {
	var c = Command{}
	var err = r.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (at *Command) ParseCommand() (string, []string) {
	var words = strings.Fields(at.Text)
	var command string
	var params []string

	if len(words) > 0 {
		command = words[0]

		if len(words) > 1 {
			params = words[1:]
		}
	}

	return command, params
}
