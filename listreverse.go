package piscine

func ListReverse(l *List) {
	current := l.Head
	var previous *NodeL
	previous = nil
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
}
