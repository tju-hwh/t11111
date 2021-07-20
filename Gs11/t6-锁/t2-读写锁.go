package main

import (
	"fmt"
	"sync"
	"time"
)

var number int64
var rwlock sync.RWMutex
var w sync.WaitGroup
func main() {
	start:=time.Now()
	w.Add(1010)
	for i := 0; i < 10; i++ {
		go write()
	}
	for i := 0; i < 1000; i++ {
		go read()
	}
	w.Wait()
	nowS:=time.Since(start)
	fmt.Println("耗时：",nowS)
	fmt.Println(number)

}

func write(){
	rwlock.Lock()
	number++
	time.Sleep(20*time.Millisecond)
	rwlock.Unlock()
	w.Done()
}

func read(){
	rwlock.RLock()
	fmt.Println(number)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	w.Done()
}
