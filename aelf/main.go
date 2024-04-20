package main

import "fmt"

// 1->2->3->4->5->6->7->8->null
// 1->2->5->4->3->8->7->6->null
type Node struct {
	Val  int
	Next *Node
}

func reverseK(root *Node, k int) *Node {
	count := 0
	cur := root
	for cur != nil {
		count += 1
		cur = cur.Next
	}
	offset := count % k
	cur = root
	var pre *Node
	for i := 0; i < offset; i++ {
		pre = cur
		cur = cur.Next
	}
	fmt.Println(count, offset, count-offset, cur.Val, pre.Val)

	for i := 0; i < count-offset-1; i += k {
		node, next := reverse(cur, k)
		pre.Next = node
		pre = cur
		cur = next
	}
	return root
}

func reverse(root *Node, k int) (*Node, *Node) {
	cur := root
	var pre *Node
	for i := 0; i < k; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre, cur
}

func main() {
	root := &Node{1, &Node{2, &Node{3,
		&Node{4, &Node{5, &Node{6,
			&Node{7, &Node{8, nil}}}}}}}}
	cur := reverseK(root, 3)
	for cur != nil {
		fmt.Printf("%v, ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}
