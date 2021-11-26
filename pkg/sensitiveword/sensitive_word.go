package sensitiveword

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"os"
)

var trieRoot *trieNode

func Init() {

	// 打开文件
	file, err := os.OpenFile("./static/sensi_words.txt", os.O_RDONLY, 0600)
	if err != nil {
		zap.L().Error("打开敏感词文件出错", zap.Error(err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	trieRoot = newTrieNode(' ')
	fmt.Println(trieRoot)

	// 扫描每一行
	for scanner.Scan() {
		// 构建根节点树
		buildTree([]rune(scanner.Text()))
	}

	// 添加失败节点
	addFailNode()

	SensitiveFilter("势必傻逼是你")
}

func buildTree(row []rune) {
	temp := trieRoot
	length := 0
	for _, w := range row {

		if child, hasChild := temp.getChildNode(w); hasChild {
			temp = child
		} else {
			newNode := newTrieNode(w)
			temp.addChildNode(newNode)
			temp = newNode
		}
		length++
	}
	temp.isWordEnd = true
	temp.len = length
}

func addFailNode() {
	var queue []*trieNode
	queue = append(queue, trieRoot)
	for len(queue) != 0 {
		parent := queue[0]
		queue = queue[1:]
		var temp *trieNode

		for _, child := range parent.children {
			if parent == trieRoot {
				child.setFailNode(trieRoot)
			} else {
				temp = parent.failNode
				if temp == nil {
					child.setFailNode(trieRoot)
				} else {
					for {
						if childIn, hasChild := temp.getChildNode(child.value); hasChild {
							child.setFailNode(childIn)
							break
						}
						temp = temp.getFailNode()
					}
				}
				queue = append(queue, child)
			}
		}
	}
	fmt.Println(trieRoot)
}

func find(text []rune) map[int64]int {
	var replaceMap = make(map[int64]int)
	node := trieRoot
	for i := int64(0); i < int64(len(text)); i++ {
		r := text[i]
		for _, hasChild := node.getChildNode(r); node != nil && !hasChild; {
			node = node.getFailNode()
		}

		if node == nil {
			node = trieRoot
		} else {
			child, hasChild := node.getChildNode(r)
			if hasChild {
				node = child
			} else {
				node = nil
			}
		}

		temp := node
		for temp != nil {
			if temp.isWordEnd && temp.getValue() != ' ' {
				replaceMap[i] = temp.len
			}
			temp = temp.getFailNode()
		}
	}
	return replaceMap
}

func SensitiveFilter(content string) string {
	rc := []rune(content)
	replaceMap := find(rc)
	fmt.Println(replaceMap)

	for key, value := range replaceMap {
		for i := 0; i < value; i++ {
			rc[key-int64(value)+int64(1)+int64(i)] = 42
		}
	}

	fmt.Println(string(rc))

	return " "
}
