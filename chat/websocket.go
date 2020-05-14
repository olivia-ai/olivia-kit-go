package chat

import (
	"crypto/rand"
	"fmt"

	"github.com/gorilla/websocket"
)

// A Client will connect to the Olivia's server using the Token the Information and the Locale
type Client struct {
	Information map[string]interface{}
	Locale      string
	Token       string
	Connection  *websocket.Conn
	Channel     chan string
}

// NewClient creates a new Client by generating a random token, and setting english as the
// default langauge.
// The host is also required with a boolean, if the SSL certificate is required.
func NewClient(host string, ssl bool) (Client, error) {
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
		Locale:     "en",
		Token:      generateToken(),
		Connection: connection,
		Channel:    make(chan string),
	}, nil
}

// generateToken returns a random token of 50 characters
func generateToken() string {
	b := make([]byte, 50)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
