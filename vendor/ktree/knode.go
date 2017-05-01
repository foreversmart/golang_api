package ktree

type KNode struct {
	Content  byte            `json:"content"`
	Count    int             `json:"count"`
	Children map[byte]*KNode `json:"children"`
}

func NewKNode(content byte) *KNode {
	return &KNode{
		Content:  content,
		Count:    0,
		Children: make(map[byte]*KNode),
	}
}

func (node *KNode) AddByte(b byte) (child *KNode) {
	if value, ok := node.Children[b]; ok {
		return value
	}

	child = NewKNode(b)
	node.Children[b] = child
	node.Count++

	return
}

func (node *KNode) Find(b byte) *KNode {
	return node.Children[b]
}

func (node *KNode) SubPath() []string {
	paths := make([][]byte, 0, 5)
	for childByte, child := range node.Children {

	}




}
