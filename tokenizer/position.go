package tokenizer

import "strconv"

type position int

const eoi position = -1

func (position position) String() string {
	if position == eoi {
		return "EOI"
	}

	return "position " + strconv.Itoa(int(position))
}
