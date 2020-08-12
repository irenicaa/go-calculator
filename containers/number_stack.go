package containers

// NumberStack ...
type NumberStack []float64

// Push ...
func (stack *NumberStack) Push(number float64) {
	*stack = append(*stack, number)
}

// Pop ...
func (stack *NumberStack) Pop() (float64, bool) {
	if len(*stack) == 0 {
		return 0, false
	}

	number := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	return number, true
}
