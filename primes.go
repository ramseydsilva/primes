package main

func getPrimeFactors(end int) {
	for start := 1; start <= end; start++ {
		isPrime := true
		for i := 2; i < start/2; i++ {
			if start%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			allPrimes = append(allPrimes, start)
		}
	}
}
