package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis dial error: ", err)
		return
	}
	fmt.Println("redis conn success")
	defer conn.Close()

	_, err = conn.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println("conn do error: ", err)
	}

	r, err := redis.Int(conn.Do("Get", "abc"))
	if err != nil {
		fmt.Println(" error: ", err)
	}

	fmt.Println(r)
}
