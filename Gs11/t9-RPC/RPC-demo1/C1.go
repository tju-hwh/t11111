package RPC_demo1

import (
	"log"
	"net/http"
	"net/rpc"
)

type Arith struct {

}

type ArithRequest struct {
	A,B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func (a *Arith) Multiply(req ArithRequest,resp *ArithResponse) error {
	resp.Pro = req.A*req.B
	return nil
}

func (a *Arith) Devide(req ArithRequest,resp *ArithResponse) error {
	resp.Quo=req.A/req.B
	resp.Rem=req.A%req.B
	return nil
}

func Server() {
	//1.注册服务
	rect := new(Arith)
	//2注册一个rect服务(其实就是注册了一个类)
	rpc.Register(rect)
	//3为服务处理绑定到http协议上
	rpc.HandleHTTP()
	//4监听服务
	err:=http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}
