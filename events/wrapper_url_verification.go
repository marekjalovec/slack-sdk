package events

type WrapperUrlVerification struct {
	Type      WrapperType `json:"type"`
	Challenge string      `json:"challenge"`
	Token     string      `json:"token"` // deprecated
}
