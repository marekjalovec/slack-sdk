package blocks

// OptionObject https://api.slack.com/reference/block-kit/composition-objects#option
type OptionObject struct {
	Text        *TextObject      `json:"text"`
	Value       string           `json:"value"`
	Description *PlainTextObject `json:"description,omitempty"`
}

func NewOptionObject(text *TextObject, value string) *OptionObject {
	return &OptionObject{
		Text:  text,
		Value: value,
	}
}
