package blocks

// SelectUsersElement https://api.slack.com/reference/block-kit/block-elements#users_select
type SelectUsersElement struct {
	Type        ElementType `json:"type"`
	ActionId    string      `json:"action_id"`
	InitialUser string      `json:"initial_user,omitempty"`
	FocusOnLoad bool        `json:"focus_on_load,omitempty"`
	Placeholder *TextObject `json:"placeholder,omitempty"`
}

func (at SelectUsersElement) ElementType() ElementType {
	return at.Type
}

func NewSelectUsersElement(actionId string, initialUser string) *SelectUsersElement {
	return &SelectUsersElement{
		Type:        ElementTypeSelectUsers,
		ActionId:    actionId,
		InitialUser: initialUser,
	}
}
