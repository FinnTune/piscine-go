package piscine

func ListClear(l *List) {
	l.Head = l.Tail.Next
}
