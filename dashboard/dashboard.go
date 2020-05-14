package dashboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Client will interact with the REST Api routes of Olivia
type Client struct {
	URL    string
	Token  string
	Locale string
}

// Data contains the data sent for the dashboard
type Data struct {
	Layers   Layers   `json:"layers"`
	Training Training `json:"training"`
}

// Layers contains the data of the network's layers
type Layers struct {
	InputNodes   int `json:"input"`
	HiddenLayers int `json:"hidden"`
	OutputNodes  int `json:"output"`
}

// Training contains the data related to the training of the network
type Training struct {
	Rate   float64   `json:"rate"`
	Errors []float64 `json:"errors"`
	Time   float64   `json:"time"`
}

// GetDashboardData returns the dashboard data from the client URL of the Olivia's API
func (client *Client) GetDashboardData() (data Data, err error) {
	resp, err := http.Get(
		fmt.Sprintf("%s/api/%s/dashboard", client.URL, client.Locale),
	)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	json.Unmarshal(body, &data)

	return data, nil
}
