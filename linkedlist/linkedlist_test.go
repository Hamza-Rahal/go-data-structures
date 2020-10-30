package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArray(t *testing.T) {
	expected := 0
	ll := New()
	assert.Equal(t, expected, len(ll.ToArray()))
}

func TestAppend(t *testing.T) {
	ll := New()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	assert.Equal(t, 3, len(ll.ToArray()))
	assert.Equal(t, 1, ll.ToArray()[0])
}

func TestAddFirst(t *testing.T) {
	ll := New()
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)

	assert.Equal(t, 3, len(ll.ToArray()))
	assert.Equal(t, 3, ll.ToArray()[0])
}

func TestRemoveHead(t *testing.T) {
	ll := New()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.RemoveHead()

	assert.Equal(t, 2, len(ll.ToArray()))
	assert.Equal(t, 2, ll.ToArray()[0])
}
