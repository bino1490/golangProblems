package main

import "fmt"

type node struct {
	data int
	next *node
}

func (n *node) Addnode(d int) {
	if n.data == 0 {
		n.data = d
		return
	}
	var nxt node
	nxt.data = d
	for n.next != nil {
		n = n.next
	}
	n.next = &nxt
}

func main() {
	var n node
	n.Addnode(3)
	n.Addnode(4)
	n.Addnode(5)
	n.Addnode(6)
	n.Addnode(9)

	// Display
	for n.next != nil {
		fmt.Printf("%v ->", n.data)
		n = *n.next
	}
	fmt.Println(n.data)
}
