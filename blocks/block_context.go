package blocks

// ContextBlock https://api.slack.com/reference/block-kit/blocks#context
type ContextBlock struct {
	Type     BlockType      `json:"type"`
	Elements *[]*TextObject `json:"elements"`
	BlockId  string         `json:"block_id,omitempty"`
}

func (at ContextBlock) BlockType() BlockType {
	return at.Type
}

func NewContextBlock(elements ...*TextObject) *ContextBlock {
	return &ContextBlock{
		Type:     BlockTypeContext,
		Elements: &elements,
	}
}

func NewContextBlockWithText(text string) *ContextBlock {
	return &ContextBlock{
		Type: BlockTypeContext,
		Elements: &[]*TextObject{
			NewTextObject(TextObjectTypeMarkdown, text),
		},
	}
}
