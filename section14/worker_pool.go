package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Result struct {
	year     int
	total    int
	interest int
	id       int
}

func calc(id, price int, interestRate float64, year int, result chan Result) {
	months := year * 12
	interest := 0
	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestRate / 12)
	}
	result <- Result{year: year, total: price + interest, interest: interest, id: id}
}

func worker(id, price int, interestRate float64, years chan int, result chan Result, wg *sync.WaitGroup) {
	for year := range years {
		calc(id, price, interestRate, year, result)
		wg.Done()
	}
}

func main() {
	price := 40000000
	interestRate := 0.001
	years := make(chan int, 35)
	for i := 1; i < 36; i++ {
		years <- i
	}
	var wg sync.WaitGroup
	wg.Add(35)
	result := make(chan Result, 35)
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestRate, years, result, &wg)
	}

	wg.Wait()
	close(result)
	close(years)

	for r := range result {
		fmt.Printf("%+v\n", r)
	}
}
