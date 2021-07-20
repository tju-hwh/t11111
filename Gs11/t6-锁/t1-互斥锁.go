package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var x int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go addnumber()
	go addnumber()
	wg.Wait()
	fmt.Println(x)
}

func addnumber()  {
	for i := 0; i < 100; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}
