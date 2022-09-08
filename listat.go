package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	n := pos
	for n > 0 {
		if l.Next == nil {
			return nil
		}
		l = l.Next
		n--
	}
	return l
	/* for l != nil { // This works too counting up to pos instead of down to >0
		if n == pos {
			return l
		}
		l = l.Next
		n++
	}
	return nil */
}
