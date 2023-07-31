package blocks

// PlainTextInputElement https://api.slack.com/reference/block-kit/block-elements#input
type PlainTextInputElement struct {
	Type                 ElementType           `json:"type"`
	ActionId             string                `json:"action_id"`
	InitialValue         string                `json:"initial_value,omitempty"`
	Multiline            bool                  `json:"multiline,omitempty"`
	FocusOnLoad          bool                  `json:"focus_on_load,omitempty"`
	MinLength            int                   `json:"min_length,omitempty"`
	MaxLength            int                   `json:"max_length,omitempty"`
	Placeholder          *PlainTextObject      `json:"placeholder,omitempty"`
	DispatchActionConfig *DispatchActionConfig `json:"dispatch_action_config,omitempty"`
}

func (at PlainTextInputElement) ElementType() ElementType {
	return at.Type
}

func NewSimpleTextInputElement(actionId string, initialValue string) *PlainTextInputElement {
	return &PlainTextInputElement{
		Type:         ElementTypePlainTextInput,
		ActionId:     actionId,
		InitialValue: initialValue,
		Multiline:    false,
	}
}

func NewMultilinePlainTextInputElement(actionId string, initialValue string) *PlainTextInputElement {
	return &PlainTextInputElement{
		Type:         ElementTypePlainTextInput,
		ActionId:     actionId,
		InitialValue: initialValue,
		Multiline:    true,
	}
}
