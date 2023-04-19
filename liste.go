package main

import "fmt"

type listNode struct {
	chiave int
	next   *listNode
}
type linkedList struct {
	testa *listNode
}

func newNode(val int) *listNode {
	return &listNode{val, nil}
}
func addnewNode(l *linkedList, val int) {
	node := newNode(val)
	node.next = l.testa
	l.testa = node
}

func printList(l *linkedList) {
	p := l.testa
	for p != nil {
		fmt.Println(p.chiave, " ")
		p = p.next
	}
}

func searchList(l *linkedList, el int) bool {
	p := l.testa
	for p != nil {
		if p.chiave == el {
			return true
		}
		p = p.next
	}
	return false
}

func removeitem(l *linkedList, el int) bool {
	var prev, curr *listNode
	curr = l.testa
	prev = nil
	for curr != nil {
		if curr.chiave == el {
			if prev == nil {
				l.testa = l.testa.next
			} else {
				prev.next = curr.next
			}
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}
