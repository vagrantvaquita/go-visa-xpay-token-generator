package xpay

func NewXPayToken(tc *TokenConfig) Token {
	return tc.GenerateXPayToken()
}
