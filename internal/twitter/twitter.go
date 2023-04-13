package twitter

import (
	"internal/utils"

	"github.com/dghubble/oauth1"
)

type Twitter struct {
	API_KEY_SECRET      string
	API_KEY             string
	ACCESS_TOKEN        string
	ACCESS_TOKEN_SECRET string
}

func NewTwitter() *Twitter {

	return &Twitter{
		API_KEY_SECRET:      "",
		API_KEY:             "",
		ACCESS_TOKEN_SECRET: "",
		ACCESS_TOKEN:        "",
	}
}

func (t *Twitter) UpdateTwitterProfile(pseudo string) error {

	config := oauth1.NewConfig(t.API_KEY, t.API_KEY_SECRET)
	token := oauth1.NewToken(t.ACCESS_TOKEN, t.ACCESS_TOKEN_SECRET)

	httpClient := config.Client(oauth1.NoContext, token)

	path := utils.GeneratePath(pseudo)
	resp, err := httpClient.Post(path, "application/json;charset=utf-8", nil)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
