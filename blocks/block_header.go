package blocks

// HeaderBlock https://api.slack.com/reference/block-kit/blocks#header
type HeaderBlock struct {
	Type    BlockType        `json:"type"`
	Text    *PlainTextObject `json:"text"`
	BlockId string           `json:"block_id,omitempty"`
}

func (at HeaderBlock) BlockType() BlockType {
	return at.Type
}

func NewHeaderBlock(text string) *HeaderBlock {
	return &HeaderBlock{
		Type: BlockTypeHeader,
		Text: NewPlainTextObject(text),
	}
}
