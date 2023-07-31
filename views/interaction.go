package views

import (
	"github.com/marekjalovec/slack-sdk/blocks"
)

type Interaction struct {
	//Token               string          `json:"token"` // deprecated
	Type                InteractionType `json:"type"`
	View                *View           `json:"view,omitempty"`
	ApiAppId            string          `json:"api_app_id"`
	TriggerId           string          `json:"trigger_id"`
	IsEnterpriseInstall bool            `json:"is_enterprise_install"`
	Enterprise          string          `json:"enterprise,omitempty"`
	//CallbackId          string               `json:"callback_id,omitempty"`
	Actions   []*InteractionAction `json:"actions,omitempty"`
	Container struct {
		Type string `json:"type"`
		// type = message
		MessageTs   string `json:"message_ts,omitempty"`
		ChannelId   string `json:"channel_id,omitempty"`
		IsEphemeral bool   `json:"is_ephemeral,omitempty"`
		// type = view
		ViewId string `json:"view_id,omitempty"`
	} `json:"container,omitempty"`
	User struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Name     string `json:"name"`
		TeamId   string `json:"team_id"`
	} `json:"user"`
	Channel struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	Team struct {
		Id     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"team"`
}

type InteractionType string

const (
	InteractionTypeBlockActions   InteractionType = "block_actions"
	InteractionTypeMessageActions InteractionType = "message_actions"
	InteractionTypeViewClosed     InteractionType = "view_closed"
	InteractionTypeViewSubmission InteractionType = "view_submission"
	InteractionTypeShortcut       InteractionType = "shortcut"
)

func (at *Interaction) HasAction(actionType InteractionActionType, actionId string) bool {
	for _, action := range at.Actions {
		if action.Type == actionType && action.ActionId == actionId {
			return true
		}
	}

	return false
}

func (at *Interaction) HasButtonAction(actionId string) bool {
	return at.HasAction(InteractionActionTypeButton, actionId)
}

type InteractionActionType string

const (
	InteractionActionTypeButton         InteractionActionType = "button"
	InteractionActionTypePlainTextInput InteractionActionType = "plain_text_input"
	InteractionActionTypeStaticSelect   InteractionActionType = "static_select"
	InteractionActionTypeUrlTextInput   InteractionActionType = "url_text_input"
	InteractionActionTypeUsersSelect    InteractionActionType = "users_select"
)

type InteractionAction struct {
	Type           InteractionActionType `json:"type"`
	ActionId       string                `json:"action_id,omitempty"`
	BlockId        string                `json:"block_id,omitempty"`
	ActionTs       string                `json:"action_ts,omitempty"`
	Value          string                `json:"value,omitempty"`
	SelectedUser   string                `json:"selected_user,omitempty"`
	SelectedOption *blocks.OptionObject  `json:"selected_option,omitempty"`
}
