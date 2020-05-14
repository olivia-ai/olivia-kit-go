package dashboard

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Intent is a way to group sentences that mean the same thing and link them with a tag which
// represents what they mean, some responses that the bot can reply and a context
type Intent struct {
	Tag       string   `json:"tag"`
	Patterns  []string `json:"patterns"`
	Responses []string `json:"responses"`
	Context   string   `json:"context"`
}

// GetIntents returns the array of intents for the given client from the Olivia REST API
func (client Client) GetIntents() (intents []Intent, err error) {
	resp, err := http.Get(
		fmt.Sprintf("%s/api/%s/intents", client.URL, client.Locale),
	)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	json.Unmarshal(body, &intents)

	return intents, nil
}

// CreateIntents creates the given intent in the given Olivia REST API and returns an error
func (client Client) CreateIntents(intent Intent) error {
	b, err := json.Marshal(intent)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/api/%s/intent", client.URL, client.Locale),
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response map[string]interface{}
	json.Unmarshal(body, &response)

	message, found := response["message"]
	if found {
		return errors.New(string(fmt.Sprintf("%v", message)))
	}

	return nil
}
