package calculator

import "strconv"

type position int

const eoi position = -1

func (position position) String() string {
	if position == eoi {
		return "EOI"
	}

	return strconv.Itoa(int(position))
}
