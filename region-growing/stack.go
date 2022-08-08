package regiongrowing

import "github.com/simonmau/spacial-base-calculation/point"

type Stack []point.T

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str *point.T) {
	*s = append(*s, *str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (point *point.T, ok bool) {
	if s.IsEmpty() {
		return nil, false
	}

	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return &element, true
}
