package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	nwNod := &NodeI{}
	nwNod.Data = data_ref
	nwNod.Next = nil

	if l == nil || l.Data >= nwNod.Data {
		nwNod.Next = l
		return nwNod
	} else {
		temporary := l
		for temporary.Next != nil && temporary.Next.Data < nwNod.Data {
			temporary = temporary.Next
		}
		nwNod.Next = temporary.Next
		temporary.Next = nwNod
	}
	return l
}
