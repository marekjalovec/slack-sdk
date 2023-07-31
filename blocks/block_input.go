package blocks

// InputBlock https://api.slack.com/reference/block-kit/blocks#input
type InputBlock struct {
	Type           BlockType        `json:"type"`
	Label          *PlainTextObject `json:"label,omitempty"`
	Element        *ElementWrapper  `json:"element,omitempty"`
	Hint           *PlainTextObject `json:"hint,omitempty"`
	DispatchAction bool             `json:"dispatch_action,omitempty"`
	Optional       bool             `json:"optional,omitempty"`
	BlockId        string           `json:"block_id,omitempty"`
}

func (at InputBlock) BlockType() BlockType {
	return at.Type
}

func NewInputBlock(element Element, label string) *InputBlock {
	var e *ElementWrapper
	if element != nil {
		e = NewElementWrapper(element)
	}

	return &InputBlock{
		Type:    BlockTypeInput,
		Label:   NewPlainTextObject(label),
		Element: e,
	}
}
