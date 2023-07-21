package blocks

// ActionsBlock https://api.slack.com/reference/block-kit/blocks#actions
type ActionsBlock struct {
	Type     BlockType         `json:"type"`
	Elements []*ElementWrapper `json:"elements"`
	BlockId  string            `json:"block_id,omitempty"`
}

func (at ActionsBlock) BlockType() BlockType {
	return at.Type
}

func NewActionsBlock(elements ...Element) *ActionsBlock {
	var e []*ElementWrapper
	for _, element := range elements {
		e = append(e, NewElementWrapper(element))
	}

	return &ActionsBlock{
		Type:     BlockTypeActions,
		Elements: e,
	}
}
