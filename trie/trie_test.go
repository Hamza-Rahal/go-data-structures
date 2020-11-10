package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	trie := New()
	trie.Insert([]byte("celestial"))

	b := trie.Contains([]byte("celestial"))
	assert.Equal(t, true, b)

	b = trie.Contains([]byte("truth"))
	assert.Equal(t, false, b)
}
