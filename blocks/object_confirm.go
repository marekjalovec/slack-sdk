package blocks

// ConfirmObject https://api.slack.com/reference/block-kit/composition-objects#confirm
type ConfirmObject struct {
	Title   *PlainTextObject   `json:"title"`
	Text    *TextObject        `json:"text"`
	Confirm *PlainTextObject   `json:"confirm"`
	Deny    *PlainTextObject   `json:"deny"`
	Style   ButtonElementStyle `json:"style,omitempty"`
}

func NewConfirmObject(title string, text *TextObject, confirm string, deny string) *ConfirmObject {
	return &ConfirmObject{
		Title:   NewPlainTextObject(title),
		Text:    text,
		Confirm: NewPlainTextObject(confirm),
		Deny:    NewPlainTextObject(deny),
	}
}
