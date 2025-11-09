package main

import "github.com/julian-richter/leetcode/AddTwoNumbers/types"

/*
okay, so here's the situation:

we’ve got two linked lists, both representing numbers, but the digits are reversed
like [2 -> 4 -> 3] actually means 342, and [5 -> 6 -> 4] means 465

we need to add them and return a new linked list of the sum
so for 342 + 465 = 807
our output linked list should be [7 -> 0 -> 8]

each node is literally one digit, and they’re in reverse order
which is honestly a blessing because it means we can just walk through both lists left to right
and handle carry like normal addition

the plan:
- use a dummy node to avoid crying about head initialization
- keep track of a carry for when sums go over 9
- move through both lists at the same time
- build a new list as we go
*/

func addTwoNumbers(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	// the dummy node exists purely so we don’t have to special-case the head
	dummy := &types.ListNode{}
	currentPointer := dummy
	carry := 0 // if we go over 9, we’ll carry 1 into the next round like in elementary school

	// keep looping while there’s *something* to add
	// could be leftover digits in l1, l2, or a sneaky carry that won’t go away
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry // start with the carry from last time, if there was one

		// if l1 still has digits left, add them and move on
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// same thing for l2
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// if sum goes double digits, carry the 1 for next time
		carry = sum / 10

		// only keep the one's place in this node
		// (the % 10 basically chops off the tens digit)
		currentPointer.Next = &types.ListNode{Val: sum % 10}

		// move the pointer forward
		currentPointer = currentPointer.Next
	}

	// return the actual head, dummy’s done
	return dummy.Next
}
