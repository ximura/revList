package list

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestList(t *testing.T) {
	t.Parallel()

	t.Run("Push to empty list", func(t *testing.T) {
		var l LinkedList[int]
		l.Push(1)

		assert.Assert(t, l.head != nil, "Head is empty")
		assert.Assert(t, l.head.Value == 1, "Head value is incorrect")
		assert.Assert(t, l.head.Next == nil, "Head next must  be empty")
	})

	t.Run("Push to list", func(t *testing.T) {
		var l LinkedList[int]
		l.Push(1)
		l.Push(2)

		assert.Assert(t, l.head != nil, "Head is empty")
		assert.Assert(t, l.head.Value == 1, "Head value is incorrect")
		assert.Assert(t, l.head.Next != nil, "Head next must be not empty")
		assert.Assert(t, l.head.Next.Value == 2)
		assert.Assert(t, l.head.Next.Next == nil)
	})

	t.Run("Reverse empty", func(t *testing.T) {
		var l LinkedList[int]
		l.Reverse()

		assert.Assert(t, l.head == nil, "Head is not empty")
	})

	t.Run("Reverse only head", func(t *testing.T) {
		var l LinkedList[int]
		l.Push(1)
		l.Reverse()

		assert.Assert(t, l.head != nil, "Head is empty")
		assert.Assert(t, l.head.Value == 1, "Head value is incorrect")
		assert.Assert(t, l.head.Next == nil, "Head next must  be empty")
	})

	t.Run("Reverse 1->2->3 into 3->2->1", func(t *testing.T) {
		var l LinkedList[int]
		l.Push(1)
		l.Push(2)
		l.Push(3)
		l.Reverse()

		assert.Assert(t, l.head != nil, "Head is empty")
		assert.Equal(t, l.head.Value, 3)
		assert.Equal(t, l.head.Next.Value, 2)
		assert.Equal(t, l.head.Next.Next.Value, 1)
		assert.Assert(t, l.head.Next.Next.Next == nil)
	})

	t.Run("Stringer empty", func(t *testing.T) {
		var l LinkedList[int]

		s := l.String()
		assert.Equal(t, s, "", "List string representation is incorrect")
	})

	t.Run("Stringer", func(t *testing.T) {
		var l LinkedList[int]
		for i := 0; i < 10; i++ {
			l.Push(i)
		}

		s := l.String()
		assert.Equal(t, s, "0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9", "List string representation is incorrect")
	})

	t.Run("Stringer Reverse", func(t *testing.T) {
		var l LinkedList[int]
		for i := 0; i < 10; i++ {
			l.Push(i)
		}

		l.Reverse()
		s := l.String()
		assert.Equal(t, s, "9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1 -> 0", "List string representation is incorrect")
	})
}
