package events

import (
	slacksdk "github.com/marekjalovec/slack-sdk/client"
)

type WrapperType string

const (
	WrapperTypeUrlVerification WrapperType = "url_verification"
	WrapperTypeEventCallback   WrapperType = "event_callback"
)

func GetWrapperType(r *slacksdk.Request) (WrapperType, error) {
	var t = struct {
		Type WrapperType `json:"type"`
	}{}
	err := r.Unmarshal(&t)
	if err != nil {
		return "", err
	}

	return t.Type, nil
}
