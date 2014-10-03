package main

import (
	"log"
	"os"
	"strconv"
)

type Factors []int

type Number struct {
	number       int
	factors      *Factors
	primeFactors *Factors
	isPrime      bool
}

var allPrimes = []int{}
var results = map[int]Number{}

func main() {
	if len(os.Args) > 1 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		} else {
			if len(os.Args) > 2 {
				getPrimeFactors(num)
			} else {
				num := getFactors(num)
				log.Println(num.factors, num.primeFactors)
			}
			log.Println(allPrimes)
		}
	}
}
