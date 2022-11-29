package stack

// Stack for elements.
type Stack[T any] struct {
	top  *Element[T]
	size int
}

// Element is the single element of the Stack.
type Element[T any] struct {
	value T
	next  *Element[T]
}

// Len returns the stack's length.
func (s *Stack[T]) Len() int {
	return s.size
}

// Push a new element onto the stack.
func (s *Stack[T]) Push(value T) {
	s.top = &Element[T]{value, s.top}
	s.size++
}

// Top returns a value of the top element of the stack.
func (s *Stack[T]) Top() (value T, ok bool) {
	if s.size > 0 {
		return s.top.value, true
	}

	return
}

// Pop removes the top element from the stack and returns it's value.
func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--

		return value, true
	}

	return
}
