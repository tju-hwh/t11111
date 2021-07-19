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

	_, err = conn.Do("MSet", "abc", 100,"ccc",12)
	if err != nil {
		fmt.Println("conn do error: ", err)
	}

	r, err := redis.Ints(conn.Do("MGet", "abc","ccc"))
	if err != nil {
		fmt.Println(" error: ", err)
	}

	for k, v := range r {
		fmt.Println(k,v)
	}
}
