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
		API_KEY_SECRET:      "swPe20Z7mVWskHInyXX4iuHygVboK4RXxPoiKzPTFMsruAYizE",
		API_KEY:             "XbH4MpafUifzc0IyY3KYosOOw",
		ACCESS_TOKEN_SECRET: "uJWlcbDG35rPTk4fodPRlLnJ1w89fyzgMMGPlfE3m4Hhe",
		ACCESS_TOKEN:        "872417601846218752-gBeOuJXS2BIm3tEgIXJdPPXTpjNfcJ4",
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
