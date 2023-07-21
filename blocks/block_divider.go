package blocks

// DividerBlock https://api.slack.com/reference/block-kit/blocks#divider
type DividerBlock struct {
	Type    BlockType `json:"type"`
	BlockId string    `json:"block_id,omitempty"`
}

func (at DividerBlock) BlockType() BlockType {
	return at.Type
}

func NewDividerBlock() *DividerBlock {
	return &DividerBlock{
		Type: BlockTypeDivider,
	}
}
