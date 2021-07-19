package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init(){
	pool=&redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",":6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         16,
		MaxActive:       0,
		IdleTimeout:     300,
	}
}

func main() {

	c:=pool.Get()
	defer c.Close()

	_, err := c.Do("Set", "zz", 10)
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	i, err := redis.Int(c.Do("Get", "zz"))
	if err != nil {
		fmt.Println(" error: ", err)
		return
	}
	fmt.Println(i)
	pool.Close()
}

