package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	count := 0
	var initial *NodeI

	if l == nil || l.Next == nil {
		return l
	}

	for l != nil {
		next := l.Next
		if count == 0 {
			initial = l
			count++
		}

		for next != nil {
			if l.Data > next.Data {
				l.Data, next.Data = next.Data, l.Data
			}
			next = next.Next
		}
		l = l.Next
	}
	return initial
}
