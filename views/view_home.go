package views

import (
	"github.com/marekjalovec/slack-sdk/blocks"
	"github.com/marekjalovec/slack-sdk/client"
)

type HomeViewPayload struct {
	View      *HomeView `json:"view"`
	UserId    string    `json:"user_id,omitempty"`
	ViewId    string    `json:"view_id,omitempty"`
	TriggerId string    `json:"trigger_id,omitempty"`
}

func NewHomeViewPayload(view *HomeView, userId string, viewId string) *HomeViewPayload {
	return &HomeViewPayload{
		View:   view,
		UserId: userId,
		ViewId: viewId,
	}
}

type HomeView struct {
	Id     string `json:"id,omitempty"`
	TeamId string `json:"team_id,omitempty"`
	State  *State `json:"state,omitempty"`

	Type            ViewType              `json:"type"`
	Blocks          *blocks.BlocksWrapper `json:"blocks"`
	PrivateMetadata string                `json:"private_metadata,omitempty"`
	CallbackId      string                `json:"callback_id,omitempty"`
	ExternalId      string                `json:"external_id,omitempty"`
}

func NewHomeView(b ...blocks.Block) *HomeView {
	return &HomeView{
		Type:   ViewTypeHome,
		Blocks: blocks.NewBlocksWrapper(b...),
	}
}

func (at *HomeView) Publish(client *client.Client, userId string) error {
	_, err := client.Post("views.publish", NewHomeViewPayload(at, userId, at.Id))
	return err
}
