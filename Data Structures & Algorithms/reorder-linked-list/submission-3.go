/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil

	var prev *ListNode
	for second != nil {
		next := second.Next
		second.Next = prev
		prev = second
		second = next
	}
	second = prev

	first := head
	for second != nil {
		nextFirst := first.Next
		nextSecond := second.Next

		first.Next = second
		second.Next = nextFirst

		first = nextFirst
		second = nextSecond
	}
}

