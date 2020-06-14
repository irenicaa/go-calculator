package calculator

// TokenStack ...
type TokenStack []Token

// Push ...
func (stack *TokenStack) Push(token Token) {
	*stack = append(*stack, token)
}
