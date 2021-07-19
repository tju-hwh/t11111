package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	dial, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("dial error: ", err)
		return
	}
	_, err = dial.Do("expire", "abc", 5)
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}

}
