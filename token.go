package xpay

// NewXPayToken constructs a new X-Pay Token which can be use to make
// requests to Visa API endpoints.
func NewXPayToken(tc *TokenConfig) Token {
	return tc.GenerateXPayToken()
}
