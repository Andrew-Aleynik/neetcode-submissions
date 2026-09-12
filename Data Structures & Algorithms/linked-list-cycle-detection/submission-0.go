/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func hasCycle(head *ListNode) bool {
	i, j := head, head
	for i != nil && j != nil && j.Next != nil {
		i = i.Next
		j = j.Next.Next
		if i == j {
			return true
		}
	}
	return false
}
