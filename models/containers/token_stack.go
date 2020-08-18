package containers

import "github.com/irenicaa/go-calculator/models"

// TokenStack ...
type TokenStack []models.Token

// Push ...
func (stack *TokenStack) Push(token models.Token) {
	*stack = append(*stack, token)
}

// Pop ...
func (stack *TokenStack) Pop() (models.Token, bool) {
	if len(*stack) == 0 {
		return models.Token{}, false
	}

	token := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return token, true
}
