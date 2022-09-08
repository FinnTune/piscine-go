package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	counter := l.Head
	for counter != nil {
		if comp(counter.Data, ref) {
			return &counter.Data
		}

		counter = counter.Next
	}
	return nil
}
