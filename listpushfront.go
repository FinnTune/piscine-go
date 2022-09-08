package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := NodeL{Data: data}
	/* if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
	} else {
		l.Tail.Next = &newNode
		l.Tail = &newNode
	} */
	/* 	if l.Head == nil {			// This works forwards
		l.Head = &newNode
	} else if l.Tail == nil {
		l.Head.Next = &newNode		// Redundant
		l.Tail = &newNode
	} else {
		l.Tail.Next = &newNode
		l.Tail = &newNode
	} */
	if l.Head != nil { // This Prints out backwards
		newNode.Next = l.Head
		l.Head = &newNode
	} else {
		l.Head = &newNode
	}
}
