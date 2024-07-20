package xpay

import (
	"net/url"
	"os"
	"testing"
	"time"
)

var (
	APIKey       = os.Getenv("VISA_API_KEY")
	SharedSecret = os.Getenv("VISA_SHARED_SECRET")
)

func TestXPayToken(t *testing.T) {

	query := make(url.Values)
	query.Set("apiKey=", APIKey)

	tokenConfig := &TokenConfig{
		SharedSecret: SharedSecret,
		DateTime:     time.Now(),
		ResourcePath: "helloworld",
		QueryParams:  query,
		PostBody:     "",
	}

	token := NewXPayToken(tokenConfig)
	t.Log(token)
}
