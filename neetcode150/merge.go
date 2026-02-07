// Time Complexity: O(n+m)
// Space Complexity: O(n+m)
// n is the length of list1 and m is the length of list2
// The analogy below uses a recursive approach with base
// cases of returning list2 if list1 is empty, and returning
// list1 if list2 is empty, this works when either of the list
// is nil during the recursive process, the other list elements
// is returned
// The recursive process takes a iterative apporach such that when the
// head value of list1 is smaller or equal to head value of list2, the head
// value of list1 Next pointer is to merge the two lists(remaining elements
// of list1 and list2 elements), otherwise if the head value of list2 is less than
// head value of list1, the head value of list2 Next pointer is to merge the
// two lists(list1 elements and remaining elements of list2)
package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
