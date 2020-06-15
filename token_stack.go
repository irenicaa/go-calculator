package calculator

// TokenStack ...
type TokenStack []Token

// Push ...
func (stack *TokenStack) Push(token Token) {
	*stack = append(*stack, token)
}

// Pop ...
func (stack *TokenStack) Pop() (Token, bool) {
	if len(*stack) == 0 {
		return Token{}, false
	}

	token := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return token, true
}
