package events

import (
	slacksdk "github.com/marekjalovec/slack-sdk/views"
)

type Event interface {
	GetType() EventType
}

type EventType string

const (
	EventTypeAppHomeOpened EventType = "app_home_opened"
	EventTypeLinkShared    EventType = "link_shared"
)

type EventLinkShared struct {
	Type      EventType `json:"type"`
	User      string    `json:"user"`
	Channel   string    `json:"channel"`
	MessageTs string    `json:"message_ts"`
	Links     []struct {
		Url    string `json:"url"`
		Domain string `json:"domain"`
	} `json:"links"`
	Source          string `json:"source"`
	UnfurlId        string `json:"unfurl_id"`
	IsBotUserMember bool   `json:"is_bot_user_member"`
	EventTs         string `json:"event_ts"`
}

func (at EventLinkShared) GetType() EventType {
	return at.Type
}

type EventAppHomeOpened struct {
	Type      EventType     `json:"type"`
	UserId    string        `json:"user"`
	ChannelId string        `json:"channel"`
	Tab       string        `json:"tab"` // home, messages
	View      slacksdk.View `json:"view"`
	EventTs   string        `json:"event_ts"`
}

func (at EventAppHomeOpened) GetType() EventType {
	return at.Type
}
