package node

// test
type ListNode struct {
	Val int
	// WARNING
	Next *ListNode
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}
