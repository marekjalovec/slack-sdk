package blocks

// OptionObject https://api.slack.com/reference/block-kit/composition-objects#option
type OptionObjectPlainText struct {
	Text  *PlainTextObject `json:"text"`
	Value string           `json:"value"`
}

func NewOptionObjectPlainText(text string, value string) *OptionObjectPlainText {
	return &OptionObjectPlainText{
		Text:  NewPlainTextObject(text),
		Value: value,
	}
}
