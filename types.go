package xpay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Aliasing string to use for Token return.
type Token string

// TokenConfig is the type of constructor options for NewXPayToken(...).
type TokenConfig struct {
	SharedSecret string
	DateTime     time.Time
	ResourcePath string
	QueryParams  url.Values
	PostBody     string
}

// Timestamp parses time.Time type into a string representix Unix timestamp.
func (tc TokenConfig) Timestamp() string {
	return strconv.Itoa(int(tc.DateTime.Unix()))
}

// GetMessage return a string containing each element of the message that is
// digested later as part of the specification.
func (tc TokenConfig) GetMessage() string {
	return strings.Join(
		[]string{
			tc.Timestamp(),
			tc.ResourcePath,
			tc.QueryParams.Encode(),
			tc.PostBody},
		"",
	)
}

// HashString digest the message using sha256 and hmac.
func (tc TokenConfig) HashString() string {
	h := hmac.New(sha256.New, []byte(tc.SharedSecret))
	h.Write([]byte(tc.GetMessage()))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateXPayToken each element on the token together and returns
// a string of type Token.
func (tc *TokenConfig) GenerateXPayToken() Token {
	token := strings.Join(
		[]string{
			"xv2",
			tc.Timestamp(),
			tc.HashString()},
		":",
	)
	return Token(token)
}
