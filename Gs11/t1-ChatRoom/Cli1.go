package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("dial error: ", err)
	}
	defer conn.Close()
	str:="hhhh"
	buf:=make([]byte,1024)
	write ,err:=conn.Write([]byte(str))
	if err != nil {
		fmt.Println(" error: ", err)
	}
	fmt.Println(write)
	read, err := conn.Read(buf)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	fmt.Println("读取到：",read,"内容",string(buf[0:read]))
}
