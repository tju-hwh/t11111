package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp",":8080")
	if err != nil {
		fmt.Println("net.Listen error: ", err)
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("listen accept error: ", err)
	}
	defer conn.Close()
	buf:=make([]byte,1024)
	read, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn read error: ", err)
	}
	st:= fmt.Sprintf("%s,WORLD\n",string(buf[0:read-1]))
	write, err := conn.Write([]byte(st))
	if err != nil {
		fmt.Println("write back error: ", err)
	}
	fmt.Println("写回",write)
	fmt.Println("服务段已经全部写回 ")
}
