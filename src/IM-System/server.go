package main

import (
	"fmt"
	"net"
)


type Server struct {
	Ip string
	Port int
}

//创建一个server的API
func NewServer(ip string, port int) *Server{
	server := &Server{
		Ip : ip,
		Port : port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//当前连接的业务
	fmt.Println("连接建立成功")
}

//启动服务器的API
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",this.Ip, this.Port))
	if err != nil {
		fmt.Println("监听状态启动失败",err)
		return
	}

	//close listen socket
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("某个客户端进来，但是连接失败", err)
			continue //继续等客户端连接
		}

		//如果成功，则执行当前连接成功后的业务。do handler 回调
		go this.Handler(conn)
	}
	
}