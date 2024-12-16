package solution202411

import (
	"fmt"
	"strconv"
	"strings"
)

type Cache map[string]*Node

var cache = Cache{}

type Node struct {
	Value    string
	Children []*Node
}

func (node *Node) Blink(times int) {
	if times < 1 {
		return
	}

	cachedNode, ok := cache[node.Value]
	if ok {
		fmt.Printf("Cache hit. Value: %s\n", node.Value)
		node.Children = cachedNode.Children
		return
	}

	cache[node.Value] = node

	node.Children = make([]*Node, 0, 2)

	if node.Value == "0" {
		child := &Node{
			Value: "1",
		}

		node.Children = append(node.Children, child)
	} else if len(node.Value)%2 == 0 {
		left := node.Value[:(len(node.Value) / 2)]
		right := node.Value[(len(node.Value) / 2):]

		for len(right) > 1 && strings.HasPrefix(right, "0") {
			right = strings.TrimPrefix(right, "0")
		}

		leftChild := &Node{
			Value: left,
		}
		rightChild := &Node{
			Value: right,
		}

		node.Children = append(node.Children, leftChild, rightChild)
	} else {
		value, _ := strconv.Atoi(node.Value)
		value *= 2024

		child := &Node{
			Value: strconv.Itoa(value),
		}

		node.Children = append(node.Children, child)
	}

	for _, child := range node.Children {
		child.Blink(times - 1)
	}
}

func (node *Node) ChildrenAtLevel(level int) int {
	if level == 0 {
		return 1
	}

	ret := 0
	for _, child := range node.Children {
		ret += child.ChildrenAtLevel(level - 1)
	}

	return ret
}
