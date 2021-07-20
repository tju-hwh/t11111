package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var a int64
var lock3 sync.Mutex
var wg3 sync.WaitGroup
func add(){
	a++
	wg3.Done();
}

func mutexAdd(){
	lock3.Lock()
	a++
	lock3.Unlock()
	wg3.Done()
}

func atomicAdd(){
	atomic.AddInt64(&a,1)
	wg3.Done()
}

func main() {
	t1:=time.Now()
	wg3.Add(10000)
	for i := 0; i < 10000; i++ {
		//go add()
		//go mutexAdd()
		go atomicAdd()
	}
	wg3.Wait()
	t2:=time.Now()
	t3:=t2.Sub(t1)
	fmt.Println("a: ",a," time: ",t3)
}
