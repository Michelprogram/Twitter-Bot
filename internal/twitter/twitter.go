package twitter

import (
	"errors"
	"internal/utils"
	"os"

	"github.com/dghubble/oauth1"
)

type Twitter struct {
	API_KEY_SECRET      string
	API_KEY             string
	ACCESS_TOKEN        string
	ACCESS_TOKEN_SECRET string
}

func NewTwitter() (*Twitter, error) {

	apiKeySecret := os.Getenv("API_KEY_SECRET")
	apiKey := os.Getenv("API_KEY")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")

	if apiKeySecret == "" {
		return nil, errors.New("API_KEY_SECRET not found")
	}

	if apiKey == "" {
		return nil, errors.New("API_KEY not found")
	}

	if accessTokenSecret == "" {
		return nil, errors.New("ACCESS_TOKEN_SECRET not found")
	}

	if accessToken == "" {
		return nil, errors.New("ACCESS_TOKEN not found")
	}

	return &Twitter{
		API_KEY_SECRET:      apiKeySecret,
		API_KEY:             apiKey,
		ACCESS_TOKEN_SECRET: accessTokenSecret,
		ACCESS_TOKEN:        accessToken,
	}, nil
}

func (t *Twitter) UpdateTwitterProfile(pseudo string) error {

	config := oauth1.NewConfig(t.API_KEY, t.API_KEY_SECRET)
	token := oauth1.NewToken(t.ACCESS_TOKEN, t.ACCESS_TOKEN_SECRET)

	httpClient := config.Client(oauth1.NoContext, token)

	path := utils.GeneratePath(pseudo)
	resp, err := httpClient.Post(path, "application/json;charset=utf-8", nil)

	if resp.StatusCode != 200 {
		return errors.New("Unvalid response from api")
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
