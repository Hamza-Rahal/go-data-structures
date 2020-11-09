package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	bn := New()

	assert.Equal(t, 0, bn.Size())
}

func TestInsert(t *testing.T) {
	bn := New()

	bn.Insert(1)
	bn.Insert(2)
	bn.Insert(3)
	bn.Insert(4)

	assert.Equal(t, 4, bn.Size())
}

func TestHeight(t *testing.T) {
	bn := New()

	bn.Insert(1)
	bn.Insert(1)
	bn.Insert(1)
	bn.Insert(1)

	assert.Equal(t, 2, bn.Height())
}
