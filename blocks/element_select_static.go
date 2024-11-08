package blocks

// SelectStaticElement https://api.slack.com/reference/block-kit/block-elements#static_select
type SelectStaticElement struct {
	Type          ElementType               `json:"type"`
	ActionId      string                    `json:"action_id"`
	Options       *[]*OptionObjectPlainText `json:"options,omitempty"`
	InitialOption *OptionObjectPlainText    `json:"initial_option,omitempty"`
	Confirm       *ConfirmObject            `json:"confirm,omitempty"`
	FocusOnLoad   bool                      `json:"focus_on_load,omitempty"`
	Placeholder   *PlainTextObject          `json:"placeholder,omitempty"`
}

func (at SelectStaticElement) ElementType() ElementType {
	return at.Type
}

func NewSelectStaticElement(actionId string, options *[]*OptionObjectPlainText) *SelectStaticElement {
	return &SelectStaticElement{
		Type:     ElementTypeStaticSelect,
		ActionId: actionId,
		Options:  options,
	}
}
