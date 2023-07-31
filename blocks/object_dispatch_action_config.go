package blocks

import (
	"encoding/json"
)

// DispatchActionConfig https://api.slack.com/reference/block-kit/composition-objects#dispatch_action_config
type DispatchActionConfig struct {
	OnEnterPressed     bool
	OnCharacterEntered bool
}

const (
	dispatchActionOnEnterPressed     = "on_enter_pressed"
	dispatchActionOnCharacterEntered = "on_character_entered"
)

func NewDispatchActionConfig(onEnter bool, onCharacter bool) *DispatchActionConfig {
	return &DispatchActionConfig{
		OnEnterPressed:     onEnter,
		OnCharacterEntered: onCharacter,
	}
}

func (at *DispatchActionConfig) MarshalJSON() ([]byte, error) {
	var dac struct {
		TriggerActionsOn []string `json:"trigger_actions_on"`
	}
	if at.OnEnterPressed {
		dac.TriggerActionsOn = append(dac.TriggerActionsOn, dispatchActionOnEnterPressed)
	}
	if at.OnCharacterEntered {
		dac.TriggerActionsOn = append(dac.TriggerActionsOn, dispatchActionOnCharacterEntered)
	}

	return json.Marshal(dac)
}

func (at *DispatchActionConfig) UnmarshalJSON(data []byte) error {
	var dac struct {
		TriggerActionsOn []string `json:"trigger_actions_on"`
	}
	err := json.Unmarshal(data, &dac)
	if err != nil {
		return err
	}

	for _, v := range dac.TriggerActionsOn {
		switch v {
		case dispatchActionOnEnterPressed:
			at.OnEnterPressed = true
		case dispatchActionOnCharacterEntered:
			at.OnCharacterEntered = true
		}
	}

	return nil
}
