package ktree

import (
	"testing"

	"github.com/golib/assert"
)

func TestKTree_AddWord(t *testing.T) {
	assertion := assert.New(t)
	word := RandString(5)
	tree := NewKTree()
	tree.AddWord(word)

	node := tree.Root
	for _, b := range []byte(word) {
		value, ok := node.Children[b]
		assertion.True(ok)
		node = value
	}
}

func BenchmarkKTree_AddSameWord(b *testing.B) {
	tree := NewKTree()
	word := RandString(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.AddWord(word)
	}
}

func BenchmarkKTree_AddDifferentWord(b *testing.B) {
	tree := NewKTree()
	words := make([]string, 60000)
	for i, _ := range words {
		words[i] = RandString(100)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.AddWord(words[i%60000])
	}
}
