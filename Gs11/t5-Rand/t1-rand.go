package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a:=rand.Int()
		fmt.Println(a)
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>")
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		b:=rand.Int()
		fmt.Println(b)
	}
	fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}