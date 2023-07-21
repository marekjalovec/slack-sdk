package blocks

func NewMarkdownBlock(text string) *SectionBlock {
	return &SectionBlock{
		Type: BlockTypeSection,
		Text: NewTextObjectMarkdown(text),
	}
}
