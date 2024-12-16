package solution202411

import (
	"fmt"
	"strconv"
	"strings"
)

type Node1 struct {
	Value string
	Next  *Node1
}

func (node Node1) String() (ret string) {
	ret += node.Value

	if node.Next != nil {
		ret += " --> "
		ret += fmt.Sprint(node.Next)
	} else {
		ret += "\n"
	}

	return
}

func (node *Node1) Blink(times int) {
	if times < 1 {
		return
	}

	if node.Value == "0" {
		node.Value = "1"

		node.Blink(times - 1)
	} else if len(node.Value)%2 == 0 {
		left := node.Value[:(len(node.Value) / 2)]
		right := node.Value[(len(node.Value) / 2):]

		for len(right) > 1 && strings.HasPrefix(right, "0") {
			right = strings.TrimPrefix(right, "0")
		}

		node.Value = left

		newNode1 := &Node1{
			Value: right,
		}

		if node.Next != nil {
			newNode1.Next = node.Next
		}

		node.Next = newNode1

		/*
			We must expand the next node first or keep a reference to it,
			otherwise when we call node.Next it might be a different node
		*/
		node.Next.Blink(times - 1)
		node.Blink(times - 1)
	} else {
		value, _ := strconv.Atoi(node.Value)
		value *= 2024

		node.Value = strconv.Itoa(value)

		node.Blink(times - 1)
	}
}

func (node *Node1) LengthToTail() int {
	ret := 1
	cur := node

	for cur.Next != nil {
		cur = cur.Next
		ret++
	}

	return ret
}
