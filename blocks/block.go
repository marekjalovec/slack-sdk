package blocks

import (
	"encoding/json"
)

type Block interface {
	BlockType() BlockType
}

type BlockType string

const (
	BlockTypeActions BlockType = "actions"
	BlockTypeContext BlockType = "context"
	BlockTypeDivider BlockType = "divider"
	BlockTypeHeader  BlockType = "header"
	BlockTypeInput   BlockType = "input"
	BlockTypeSection BlockType = "section"

	//BlockTypeImage    BlockType = "image"
	//BlockTypeFile     BlockType = "file"
	//BlockTypeRichText BlockType = "rich_text"
)

type BlocksWrapper struct {
	Items *[]Block
}

func (at *BlocksWrapper) AddBlock(block Block) {
	var items = append(*at.Items, block)
	at.Items = &items
}

func (at *BlocksWrapper) AddBlocks(blocks ...Block) {
	var items = append(*at.Items, blocks...)
	at.Items = &items
}

func (at *BlocksWrapper) MarshalJSON() ([]byte, error) {
	var bytes, err = json.Marshal(at.Items)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (at *BlocksWrapper) UnmarshalJSON(data []byte) error {
	var rawItems []json.RawMessage
	var err = json.Unmarshal(data, &rawItems)
	if err != nil {
		return err
	}

	var blocks []Block
	for _, item := range rawItems {
		var t = struct {
			Type BlockType `json:"type"`
		}{}
		err := json.Unmarshal(item, &t)
		if err != nil {
			return err
		}

		var block Block
		switch t.Type {
		case BlockTypeActions:
			block = &ActionsBlock{}
		case BlockTypeContext:
			block = &ContextBlock{}
		case BlockTypeDivider:
			block = &DividerBlock{}
		//case "file":
		//	block = &FileBlock{}
		case BlockTypeHeader:
			block = &HeaderBlock{}
		//case "image":
		//	block = &ImageBlock{}
		case BlockTypeInput:
			block = &InputBlock{}
		//case "rich_text":
		//	block = &RichTextBlock{}
		case BlockTypeSection:
			block = &SectionBlock{}
		default:
			panic(t.Type)
			//	block = &UnknownBlock{}
		}

		err = json.Unmarshal(item, block)
		if err != nil {
			return err
		}

		blocks = append(blocks, block)
	}

	at.Items = &blocks
	return nil
}

func NewBlocksWrapper(blocks ...Block) *BlocksWrapper {
	return &BlocksWrapper{
		Items: &blocks,
	}
}
