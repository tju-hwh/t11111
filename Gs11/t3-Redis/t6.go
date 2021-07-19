package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("lpush", "book_list", "yuwen", "shuxue", "yingyu", 300)
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	s, err := redis.String(conn.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	fmt.Println(s)
}
