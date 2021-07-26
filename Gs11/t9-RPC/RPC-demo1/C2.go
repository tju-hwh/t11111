package RPC_demo1

import (
	"fmt"
	"log"
	"net/rpc"
)

func Client(){
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
		return
	}
	req:=ArithRequest{9,2}

	var resp ArithResponse
	err = conn.Call("Arith.Multiply", req, &resp)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = conn.Call("Arith.Devide", req, &resp)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("积：%d 商：%d 余数：%d",resp.Pro,resp.Quo,resp.Rem)

}
