package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int
	B int
}
type Reply struct {
	C int
}
type Arith struct {

}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Println("计算服务：")
	return nil
}

var(
	addr = flag.String("addr","localhost:8972","server address")
)

func main() {
	flag.Parse()
	s:=server.NewServer()
	s.RegisterName("Arith",new(Arith),"")
	err:=s.Serve("tcp",*addr)
	if err != nil {
		panic(err)
	}
}