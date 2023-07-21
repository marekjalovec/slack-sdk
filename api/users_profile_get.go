package api

import (
	"encoding/json"
	"fmt"

	"github.com/marekjalovec/slack-sdk/client"
)

type userProfileResponse struct {
	Ok      bool        `json:"ok"`
	Profile UserProfile `json:"profile"`
	Error   string      `json:"error"`
}

type UserProfile struct {
	Title                   string           `json:"title"`
	Phone                   string           `json:"phone"`
	Skype                   string           `json:"skype"`
	RealName                string           `json:"real_name"`
	RealNameNormalized      string           `json:"real_name_normalized"`
	DisplayName             string           `json:"display_name"`
	DisplayNameNormalized   string           `json:"display_name_normalized"`
	Fields                  map[string]Field `json:"fields"`
	StatusText              string           `json:"status_text"`
	StatusEmoji             string           `json:"status_emoji"`
	StatusEmojiDisplayInfo  []interface{}    `json:"status_emoji_display_info"`
	StatusExpiration        int              `json:"status_expiration"`
	AvatarHash              string           `json:"avatar_hash"`
	ImageOriginal           string           `json:"image_original"`
	IsCustomImage           bool             `json:"is_custom_image"`
	Email                   string           `json:"email"`
	HuddleState             string           `json:"huddle_state"`
	HuddleStateExpirationTs int              `json:"huddle_state_expiration_ts"`
	FirstName               string           `json:"first_name"`
	LastName                string           `json:"last_name"`
	Image24                 string           `json:"image_24"`
	Image32                 string           `json:"image_32"`
	Image48                 string           `json:"image_48"`
	Image72                 string           `json:"image_72"`
	Image192                string           `json:"image_192"`
	Image512                string           `json:"image_512"`
	Image1024               string           `json:"image_1024"`
	StatusTextCanonical     string           `json:"status_text_canonical"`
}

type Field struct {
	Value string `json:"value"`
	Alt   string `json:"alt"`
}

func UsersProfileGet(c *client.Client, userId string) (*UserProfile, error) {
	body, err := c.Get("users.profile.get", map[string]string{"user": userId})
	if err != nil {
		return nil, fmt.Errorf("users.profile.get failed: %w", err)
	}

	var r = userProfileResponse{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, fmt.Errorf("JSON decode failed: %w", err)
	}

	if !r.Ok {
		if r.Error == "user_not_found" {
			return nil, nil
		}

		return nil, fmt.Errorf("users.profile.get failed: %s", r.Error)
	}

	return &r.Profile, nil
}
