package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var i int
	wg.Add(2)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		//i++
		i = 1
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		//i++
		i = 2
	}()
	wg.Wait()
	fmt.Println(i)

}
