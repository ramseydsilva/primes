package main

func getFactors(number int) *Number {
	var num Number
	var ok bool
	if num, ok = results[number]; !ok {
		num = Number{number, new(Factors), new(Factors), false}
		for i := 1; i <= number; i++ {
			if number%i == 0 {
				var t Factors
				t = append(*num.factors, i)
				num.factors = &t
				if i < number {
					if prime, ok := results[i]; ok {
						if prime.isPrime {
							t = append(*num.primeFactors, i)
							num.primeFactors = &t
						}
					} else {
						getFactors(i)
					}
				}
			}

		}
		results[number] = num
	}
	if len(*num.factors) == 2 {
		num.isPrime = true
		allPrimes = append(allPrimes, num.number)
	}
	return &num
}
