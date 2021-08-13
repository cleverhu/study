package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	isEnd    bool
	children map[string]*Node
}

func NewTree() *Tree {
	return &Tree{root: NewNode()}
}

func NewNode() *Node {
	return &Node{children: make(map[string]*Node)}
}

func (t *Tree) Add(strs ...string) {

	for _, str := range strs {
		current := t.root
		for _, s := range []rune(str) {
			if _, ok := current.children[fmt.Sprintf("%c", s)]; !ok {
				current.children[fmt.Sprintf("%c", s)] = NewNode()
			}
			current = current.children[fmt.Sprintf("%c", s)]
		}
		current.isEnd = true
	}

}

func (t *Tree) Search(str string) bool {
	current := t.root
	for _, s := range []rune(str) {
		if _, ok := current.children[fmt.Sprintf("%c", s)]; !ok {
			return false
		}
		current = current.children[fmt.Sprintf("%c", s)]
	}

	return current.isEnd == true
}

func main() {
	t := NewTree()
	t.Add("test", "gin", "gina")
	fmt.Println(t.Search("test"), t.Search("gin"), t.Search("gi"))
}
