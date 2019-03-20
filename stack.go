package stack

// Stack for elements.
type Stack struct {
	top  *Element
	size int
}

// Element is the single element of the Stack.
type Element struct {
	value interface{}
	next  *Element
}

// Len returns the stack's length.
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack.
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Top returns a pointer to the top element of the stack.
func (s *Stack) Top() interface{} {
	if s.size > 0 {
		return s.top.value
	}
	return nil
}

// Pop removes the top element from the stack and returns it's value. If the stack
// is empty nil is returned.
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
