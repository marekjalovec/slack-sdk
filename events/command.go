package events

type Command struct {
	Token               string `json:"token"`
	TeamId              string `json:"team_id"`
	TeamDomain          string `json:"team_domain"`
	ChannelId           string `json:"channel_id"`
	ChannelName         string `json:"channel_name"`
	UserId              string `json:"user_id"`
	UserName            string `json:"user_name"`
	Command             string `json:"command"`
	Text                string `json:"text"`
	ApiAppId            string `json:"api_app_id"`
	IsEnterpriseInstall string `json:"is_enterprise_install"`
	ResponseUrl         string `json:"response_url"`
	TriggerId           string `json:"trigger_id"`
}
