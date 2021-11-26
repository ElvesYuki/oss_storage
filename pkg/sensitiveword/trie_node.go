package sensitiveword

type trieNode struct {
	value     rune
	len       int
	isWordEnd bool
	children  map[rune]*trieNode
	failNode  *trieNode
}

// newTrieNode 新建节点
func newTrieNode(value rune) *trieNode {
	return &trieNode{
		value:    value,
		children: make(map[rune]*trieNode),
	}
}

// addChildNode 添加子节点
func (t *trieNode) addChildNode(child *trieNode) {
	t.children[child.value] = child
}

// getChildNode 获取子节点中指定节点
func (t *trieNode) getChildNode(key rune) (*trieNode, bool) {
	node, hasKey := t.children[key]
	if hasKey {
		return node, true
	}
	return nil, false
}

// getValue 获取节点值
func (t *trieNode) getValue() rune {
	return t.value
}

// getFailNode 获取失败节点
func (t *trieNode) getFailNode() *trieNode {
	return t.failNode
}

// setFailNode 设置失败节点
func (t *trieNode) setFailNode(failNode *trieNode) {
	t.failNode = failNode
}

// getFailNode 获取所有孩子节点
func (t *trieNode) getChildrenNode() map[rune]*trieNode {
	return t.children
}
