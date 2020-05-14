package chat

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// A Client will connect to the Olivia's server using the Token the Information and the Locale
type Client struct {
	Information *map[string]interface{}
	Locale      string
	Token       string
	Connection  *websocket.Conn
	Channel     chan string
}

// RequestMessage is the structure that uses entry connections to chat with the websocket
type RequestMessage struct {
	Type        int                    `json:"type"` // 0 for handshakes and 1 for messages
	Content     string                 `json:"content"`
	Token       string                 `json:"user_token"`
	Information map[string]interface{} `json:"information"`
	Locale      string                 `json:"locale"`
}

// NewClient creates a new Client by generating a random token, and setting english as the
// default langauge.
// You need to give a pointer of the information map of your client.
// The host is also required with a boolean, if the SSL certificate is required.
func NewClient(host string, ssl bool, information *map[string]interface{}) (Client, error) {
	// Initialite the scheme and the url
	scheme := "ws"
	if ssl {
		scheme += "s"
	}

	url := fmt.Sprintf("%s://%s", scheme, host)

	// Create the websocket connection
	connection, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return Client{}, err
	}

	defer connection.Close()

	return Client{
		Information: information,
		Locale:      "en",
		Token:       generateToken(),
		Connection:  connection,
		Channel:     make(chan string),
	}, nil
}

// Handshake performs the handshake request to the websocket
func (client *Client) Handshake() error {
	// Marshal the RequestMessage with type 0
	bytes, err := json.Marshal(RequestMessage{
		Type:        0,
		Content:     "",
		Information: *client.Information,
	})
	if err != nil {
		return err
	}

	// Write the message to the websocket
	if err = client.Connection.WriteMessage(websocket.TextMessage, bytes); err != nil {
		return err
	}

	return nil
}

// generateToken returns a random token of 50 characters
func generateToken() string {
	b := make([]byte, 50)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
