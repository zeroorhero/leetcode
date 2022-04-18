package ListNode

func removeElements(head *ListNode, val int) *ListNode {
	vir := &ListNode{
		Val:  -1,
		Next: head,
	}
	tem := vir
	for tem.Next != nil {
		if tem.Next.Val == val {
			tem.Next = tem.Next.Next
			// 注意此处应该是else的
		} else {
			tem = tem.Next
		}

	}
	return vir.Next
}
