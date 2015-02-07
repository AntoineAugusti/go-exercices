package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	// Close the channel at the end to prevent
	// deadlock
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {
	if t != nil {
		walkTree(t.Left, ch)
		ch <- t.Value
		walkTree(t.Right, ch)
	}
}

// Determine whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	values1 := make(chan int)
	values2 := make(chan int)

	go Walk(t1, values1)
	go Walk(t2, values2)

	for val := range values1 {
		if val != <-values2 {
			return false
		}
	}
	return true
}

func Show(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	Show(ch)
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
