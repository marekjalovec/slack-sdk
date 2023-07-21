package views

import (
	"github.com/marekjalovec/slack-sdk/blocks"
)

type ViewPayload struct {
	View      *View  `json:"view"`
	UserId    string `json:"user_id,omitempty"`
	ViewId    string `json:"view_id,omitempty"`
	TriggerId string `json:"trigger_id,omitempty"`
}

func NewViewPayload(view *View, userId string, viewId string) *ViewPayload {
	return &ViewPayload{
		View:   view,
		UserId: userId,
		ViewId: viewId,
	}
}

type View struct {
	Type            ViewType                `json:"type"`
	Blocks          *blocks.BlocksWrapper   `json:"blocks"`
	Id              string                  `json:"id,omitempty"`
	Title           *blocks.PlainTextObject `json:"title,omitempty"`
	Submit          *blocks.PlainTextObject `json:"submit,omitempty"`
	Close           *blocks.PlainTextObject `json:"close,omitempty"`
	PrivateMetadata string                  `json:"private_metadata,omitempty"`
	CallbackId      string                  `json:"callback_id,omitempty"`
	ClearOnClose    bool                    `json:"clear_on_close,omitempty"`
	NotifyOnClose   bool                    `json:"notify_on_close,omitempty"`
	State           *State                  `json:"state,omitempty"`
}

func NewView(viewType ViewType, blocks *blocks.BlocksWrapper) *View {
	return &View{
		Type:   viewType,
		Blocks: blocks,
	}
}

//func (at *View) AddBlock(block blocks.Block) {
//	at.Blocks.AddBlock(block)
//}
//
//func (at *View) AddBlocks(blocks ...blocks.Block) {
//	at.Blocks.AddBlocks(blocks...)
//}
