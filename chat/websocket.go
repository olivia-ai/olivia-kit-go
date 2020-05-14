package chat

import (
	"crypto/rand"
	"fmt"
)

// A Client will connect to the Olivia's server using the Token the Information and the Locale
type Client struct {
	Information map[string]interface{}
	Locale      string
	Token       string
	Channel     chan string
}

// NewClient creates a new Client by generating a random token, and setting english as the
// default langauge.
// The host is also required with a boolean, if the SSL certificate is required.
func NewClient(host string, ssl bool) Client {
	return Client{
		Locale:  "en",
		Token:   generateToken(),
		Channel: make(chan string),
	}
}

// generateToken returns a random token of 50 characters
func generateToken() string {
	b := make([]byte, 50)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
