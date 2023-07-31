package blocks

// UrlTextInputElement https://api.slack.com/reference/block-kit/block-elements#url
type UrlTextInputElement struct {
	Type                 ElementType           `json:"type"`
	ActionId             string                `json:"action_id"`
	InitialValue         string                `json:"initial_value,omitempty"`
	FocusOnLoad          bool                  `json:"focus_on_load,omitempty"`
	Placeholder          *PlainTextObject      `json:"placeholder,omitempty"`
	DispatchActionConfig *DispatchActionConfig `json:"dispatch_action_config,omitempty"`
}

func (at UrlTextInputElement) ElementType() ElementType {
	return at.Type
}

func NewUrlInputElement(actionId string, initialValue string) *UrlTextInputElement {
	return &UrlTextInputElement{
		Type:         ElementTypeUrlTextInput,
		ActionId:     actionId,
		InitialValue: initialValue,
	}
}
