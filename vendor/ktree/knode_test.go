package ktree

import (
	"testing"

	"github.com/golib/assert"
)

func TestKNode_AddByte(t *testing.T) {
	assertion := assert.New(t)

	node := NewKNode(RandByte())
	newNode := node.AddByte(RandByte())

	assertion.Equal(node.Children[newNode.Content], newNode)
}
