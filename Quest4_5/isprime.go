package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb <= 0 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		/*
			For example, consider the number 100, which is evenly divisible by these numbers:
			2, 4, 5, 10, 20, 25, 50
			Note that the largest factor, 50, is half of 100.
			This holds true for all n: all divisors are less than or equal to n/2.
		*/
		if nb%i == 0 {
			return false
		}
	}
	return true
}
