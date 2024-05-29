package main

type ListNode struct {
	val  int
	next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	for list1 != nil && list2 != nil {
		if list1.val <= list2.val {
			current.next = list1
			list1 = list1.next
		} else {
			current.next = list2
			list2 = list2.next
		}
		current = current.next
	}
	if list1 != nil {
		current.next = list1
	} else {
		current.next = list2
	}
	return current
}

func main() {
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}

	CreateLinkedlist(l1)
	CreateLinkedlist(l2)
}
