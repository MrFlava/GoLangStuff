package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(1000)
}

func main(){

	dataChan := make(chan int, 2)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	// go func(){
	// 	dataChan <- 789
	// }()

	// dataChan <- 789
	// dataChan <- 123

	// n := <- dataChan
	// fmt.Printf("n = %d\n", n)

	// n = <- dataChan
	// fmt.Printf("n = %d\n", n)

	for n := range dataChan{
		fmt.Printf("n = %d\n", n)
	}
}