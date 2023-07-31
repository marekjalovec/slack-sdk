package views

import (
	"log"
	"regexp"
	"strconv"
)

type State struct {
	Values   map[string]map[string]InteractionAction `json:"values,omitempty"`
	valueMap map[string]string
}

func (at *State) GetValues() map[string]string {
	if at.valueMap != nil {
		return at.valueMap
	}

	result := map[string]string{}
	for _, value := range at.Values {
		for actionId, action := range value {
			switch action.Type {
			case InteractionActionTypePlainTextInput:
				result[actionId] = action.Value
			case InteractionActionTypeStaticSelect:
				result[actionId] = action.SelectedOption.Value
			case InteractionActionTypeUrlTextInput:
				result[actionId] = action.Value
			case InteractionActionTypeUsersSelect:
				result[actionId] = action.SelectedUser
			default:
				panic(action.Type)
			}
		}
	}

	at.valueMap = result
	return result
}

type RegexValue struct {
	ParsedKey []string
	Value     string
}

func (at *State) GetValuesRegex(regex string) []RegexValue {
	values := at.GetValues()
	var filtered []RegexValue
	r := regexp.MustCompile(regex)

	for k, v := range values {
		m := r.FindStringSubmatch(k)
		if m == nil {
			continue
		}

		filtered = append(filtered, RegexValue{
			ParsedKey: m,
			Value:     v,
		})
	}

	return filtered
}

func (at *State) GetValue(key string) *string {
	values := at.GetValues()
	value, ok := values[key]
	if ok {
		log.Printf("[State.GetValue] unknown key %s", key)
		return &value
	}

	return nil
}

func (at *State) GetValueInt(key string) *int {
	values := at.GetValues()
	value, ok := values[key]
	if ok {
		log.Printf("[State.GetValue] unknown key %s", key)

		numValue, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("[State.GetValue] string to int conversion failed for key %s", key)
			return nil
		}

		return &numValue
	}

	return nil
}

func (at *State) ForgetValues(regex string) {
	values := map[string]map[string]InteractionAction{}
	for hash, value := range at.Values {
		for actionId, action := range value {
			match, _ := regexp.MatchString(regex, actionId)
			if !match {
				values[hash] = map[string]InteractionAction{
					actionId: action,
				}
			}
		}
	}

	at.valueMap = nil
	at.Values = values
}
