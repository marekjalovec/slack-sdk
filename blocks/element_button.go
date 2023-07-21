package blocks

// ButtonElement https://api.slack.com/reference/block-kit/block-elements#button
type ButtonElement struct {
	Type               ElementType        `json:"type"`
	Text               *PlainTextObject   `json:"text"`
	ActionId           string             `json:"action_id"`
	Url                string             `json:"url,omitempty"`
	Value              string             `json:"value,omitempty"`
	Style              ButtonElementStyle `json:"style,omitempty"`
	Confirm            *ConfirmObject     `json:"confirm,omitempty"`
	AccessibilityLabel string             `json:"accessibility_label,omitempty"`
}

type ButtonElementStyle string

const (
	ButtonElementStyleDefault ButtonElementStyle = ""
	ButtonElementStyleDanger  ButtonElementStyle = "danger"
	ButtonElementStylePrimary ButtonElementStyle = "primary"
)

func (at ButtonElement) ElementType() ElementType {
	return at.Type
}

func NewButtonElement(text string, actionId string, value string) *ButtonElement {
	return &ButtonElement{
		Type:     ElementTypeButton,
		Text:     NewPlainTextObject(text),
		ActionId: actionId,
		Value:    value,
	}
}
