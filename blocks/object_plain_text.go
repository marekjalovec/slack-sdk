package blocks

// PlainTextObject https://api.slack.com/reference/block-kit/composition-objects#text
type PlainTextObject struct {
	TextObject
}

func NewPlainTextObject(text string) *PlainTextObject {
	return &PlainTextObject{
		TextObject: TextObject{
			Type: TextObjectTypePlain,
			Text: text,
		},
	}
}
