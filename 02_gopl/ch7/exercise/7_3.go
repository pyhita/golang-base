package main

import "fmt"

// 练习 7.3： 为在gopl.io/ch4/treesort（§4.4）中的*tree类型实现一个String方法去展示tree类型的值序列。

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	// inorder
	if t == nil {
		return ""
	}

	left := t.left.String()
	right := t.right.String()

	return fmt.Sprintf("%s%d%s", left, t.value, right)
}

func main() {
	root := &tree{
		value: 2,
		left: &tree{
			value: 1,
			left:  nil,
			right: nil,
		},
		right: &tree{
			value: 3,
			left:  nil,
			right: nil,
		},
	}

	fmt.Println(root.String())
}
