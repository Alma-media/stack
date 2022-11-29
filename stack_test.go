package stack_test

import (
	"testing"

	"github.com/tiny-go/stack"
)

func TestStack(t *testing.T) {
	t.Run("Given an empty stack", func(t *testing.T) {
		stack := new(stack.Stack[int])

		t.Run("test if 'nil' is returned when Top() is called on empty stack", func(t *testing.T) {
			if _, ok := stack.Top(); ok {
				t.Error("stack was expected to be empty")
			}
		})

		stack.Push(42)
		t.Run("test if stack length is changed when element is pushed", func(t *testing.T) {
			if stack.Len() != 1 {
				t.Error("stack length was expected to be changed")
			}
		})

		top, ok := stack.Top()
		t.Run("test if Top() points to the correct element in the stack", func(t *testing.T) {
			if !ok || top != 42 {
				t.Errorf("unexpected value was returned: %v", top)
			}
		})

		pop, ok := stack.Pop()
		t.Run("test if Pop() returns expected value", func(t *testing.T) {
			if !ok || pop != 42 {
				t.Errorf("unexpected value was returned: %v", pop)
			}
		})

		t.Run("test if length is changed when the element is popped", func(t *testing.T) {
			if stack.Len() != 0 {
				t.Error("stack length was expected to be changed")
			}
		})

		t.Run("test if 'nil' is returned when Pop() is called on empty stack", func(t *testing.T) {
			if _, ok := stack.Pop(); ok {
				t.Error("stack was expected to be empty")
			}
		})
	})
}
