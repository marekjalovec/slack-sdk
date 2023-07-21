package blocks

// SectionBlock https://api.slack.com/reference/block-kit/blocks#section
type SectionBlock struct {
	Type      BlockType       `json:"type"`
	Text      *TextObject     `json:"text,omitempty"`
	Fields    *[]*TextObject  `json:"fields,omitempty"`
	Accessory *ElementWrapper `json:"accessory,omitempty"`
	BlockId   string          `json:"block_id,omitempty"`
	ActionId  string          `json:"action_id,omitempty"`
}

func (at SectionBlock) BlockType() BlockType {
	return at.Type
}

func NewSectionBlock(text *TextObject, fields *[]*TextObject, accessory Element) *SectionBlock {
	var a *ElementWrapper
	if accessory != nil {
		a = NewElementWrapper(accessory)
	}

	return &SectionBlock{
		Type:      BlockTypeSection,
		Text:      text,
		Fields:    fields,
		Accessory: a,
	}
}
