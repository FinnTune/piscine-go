package piscine

func FindNextPrime(nb int) int {
	/* if nb == 1 || nb <= 0 {
		return 0
	} */

	for !IsPrime(nb) {
		nb++
		// fmt.Printf("Trying %v\n", nb)
	}

	return nb

	/* var Prime = IsPrime(nb)
	for !Prime {
		nb++
		Prime = IsPrime(nb)
	} */

	/* else if IsPrime(nb) {
		return nb
	} else if IsPrime(nb + 1) {
		return nb + 1
	} else if IsPrime(nb + 2) {
		return nb + 2
	} */

	/* for i := nb; i <= nb-1; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true*/
}
