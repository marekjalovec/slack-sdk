package events

import (
	"encoding/json"
	"fmt"
)

type WrapperEventCallback struct {
	Type               WrapperType      `json:"type"`
	TeamId             string           `json:"team_id"`    // TODO: validate
	ApiAppId           string           `json:"api_app_id"` // TODO: validate
	Event              json.RawMessage  `json:"event"`
	EventId            string           `json:"event_id"`
	EventTime          int              `json:"event_time"`
	Authorizations     []*Authorization `json:"authorizations"`
	IsExtSharedChannel bool             `json:"is_ext_shared_channel"`
	EventContext       string           `json:"event_context"`
	Token              string           `json:"token"` // deprecated
}

type Authorization struct {
	EnterpriseId        string `json:"enterprise_id"`
	TeamId              string `json:"team_id"`
	UserId              string `json:"user_id"`
	IsBot               bool   `json:"is_bot"`
	IsEnterpriseInstall bool   `json:"is_enterprise_install"`
}

func (at *WrapperEventCallback) GetEventType() EventType {
	var t = struct {
		Type EventType `json:"type"`
	}{}
	err := json.Unmarshal(at.Event, &t)
	if err != nil {
		return ""
	}

	return t.Type
}

func (at *WrapperEventCallback) GetEvent() (Event, error) {
	et := at.GetEventType()

	var event Event
	switch et {
	case EventTypeAppHomeOpened:
		event = &EventAppHomeOpened{}
	case EventTypeLinkShared:
		event = &EventLinkShared{}
	default:
		return nil, fmt.Errorf("unknown event type: %s", et)
	}

	err := json.Unmarshal(at.Event, event)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal event: %w", err)
	}

	return event, nil
}
