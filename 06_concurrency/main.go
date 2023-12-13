package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("start")
	wg.Add(2)
	go f1(&wg)
	go f2(&wg)
	fmt.Println("End")

	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		fmt.Println("f1", i)
		i++
		if i == 50 {
			return
		}
	}
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 50; i++ {
		fmt.Println("f2", i)
	}
}
