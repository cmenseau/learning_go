package main

type node struct {
	val  int
	next *node
}

func mkList(len int) *node {
	firstNode := &node{val: 0, next: nil}

	tail := firstNode

	for i := range len {
		new_node := &node{val: i, next: nil}
		tail.next = new_node
		tail = new_node
	}

	return firstNode
}

func sumList(list *node) int {
	var tot int
	var cur *node

	for cur = list; cur.next != nil; cur = cur.next {
		tot += cur.val
	}
	tot += cur.val
	return tot
}

func mkSlice(len int) []int {
	sl := make([]int, len)

	for i := range len {
		sl[i] = i
	}

	return sl
}

func sumSlice(sl []int) int {
	tot := 0
	for i := 0; i < len(sl); i++ {
		tot += sl[i]
	}
	return tot
}
