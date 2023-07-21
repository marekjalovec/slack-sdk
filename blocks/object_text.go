package blocks

// TextObject https://api.slack.com/reference/block-kit/composition-objects#text
type TextObject struct {
	Type TextObjectType `json:"type"`
	Text string         `json:"text"`
}

type TextObjectType string

const (
	TextObjectTypePlain    TextObjectType = "plain_text"
	TextObjectTypeMarkdown TextObjectType = "mrkdwn"
)

func NewTextObject(textType TextObjectType, text string) *TextObject {
	return &TextObject{
		Type: textType,
		Text: text,
	}
}

func NewTextObjectMarkdown(text string) *TextObject {
	return NewTextObject(TextObjectTypeMarkdown, text)
}

func NewTextObjectPlain(text string) *TextObject {
	return NewTextObject(TextObjectTypePlain, text)
}
