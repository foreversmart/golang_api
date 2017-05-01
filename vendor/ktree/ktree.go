package ktree

type KTree struct {
	Root *KNode `json:"root"`
}

func NewKTree() *KTree {
	return &KTree{
		Root: NewKNode(' '),
	}
}

func (tree *KTree) AddWord(word string) {
	currentNode := tree.Root
	for _, value := range []byte(word) {
		currentNode = currentNode.AddByte(value)
	}
}

func (tree *KTree) Find(query string) *KNode {
	currentNode := tree.Root
	for _, value := range []byte(query) {
		currentNode = currentNode.Find(value)
		if currentNode == nil {
			return nil
		}
	}

	return currentNode
}
