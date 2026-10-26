// Package stack implements a basic stack
package stack

type Stack struct {
	stack []interface{}
}

// Push an element on to the stack
func (s *Stack) Push(e interface{}) {
	(*s).stack = append((*s).stack, e)
}

// Pop the first element from the stack
func (s *Stack) Pop() interface{} {
	i := s.Size() - 1
	v := (*s).stack[i]
	(*s).stack = (*s).stack[:i]
	return v
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len((*s).stack)
}

// IsEmpty returns true of the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
