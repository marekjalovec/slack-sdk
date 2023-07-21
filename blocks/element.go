package blocks

import (
	"encoding/json"
)

type Element interface {
	ElementType() ElementType
}

type ElementType string

const (
	ElementTypeButton       ElementType = "button"
	ElementTypeSelectUsers  ElementType = "users_select"
	ElementTypeStaticSelect ElementType = "static_select"
)

type ElementWrapper struct {
	Item Element
}

func (at *ElementWrapper) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(at.Item)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (at *ElementWrapper) UnmarshalJSON(data []byte) error {
	var t = struct {
		Type ElementType `json:"type"`
	}{}
	var err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}

	var item Element
	switch t.Type {
	case ElementTypeButton:
		item = &ButtonElement{}
	case ElementTypeStaticSelect:
		item = &SelectStaticElement{}
	case ElementTypeSelectUsers:
		item = &SelectUsersElement{}
	default:
		panic(t.Type)
		//	block = &UnknownBlock{}
	}

	err = json.Unmarshal(data, item)
	if err != nil {
		return err
	}

	at.Item = item
	return nil
}

func NewElementWrapper(element Element) *ElementWrapper {
	return &ElementWrapper{
		Item: element,
	}
}
