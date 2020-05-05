package tx

import (
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

// CalculateHeight returns the height of the terminal minus the diff factor
func CalculateHeight(diff uint) int {
	size, err := terminal.Height()
	if err != nil {
		size = 10
	} else if size > diff {
		size -= diff
	} else {
		size = 10
	}

	return int(size)
}
