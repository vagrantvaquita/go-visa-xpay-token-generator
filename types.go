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

type Token string

type TokenConfig struct {
	SharedSecret string
	DateTime     time.Time
	ResourcePath string
	QueryParams  url.Values
	PostBody     string
}

func (tc TokenConfig) Timestamp() string {
	return strconv.Itoa(int(tc.DateTime.Unix()))
}

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

func (tc TokenConfig) HashString() string {
	h := hmac.New(sha256.New, []byte(tc.SharedSecret))
	h.Write([]byte(tc.GetMessage()))
	return hex.EncodeToString(h.Sum(nil))
}

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
