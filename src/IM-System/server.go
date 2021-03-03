package main

import (
	"fmt"
	"net"
)


type Server struct {
	Ip string
	Port int
	
	//在线用户的列表
	OnlineMap map[string]*User
	mapLock sync.RWMutex //读写锁

	//消息广播的channel
	Message chan string

//创建一个server的接口
func NewServer(ip string, port int) *Server{
	server := &Server{
		Ip : ip,
		Port : port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string)
	}
	return server
}

//监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将msg发送给全部的在线User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}


//广播消息的方法
func(this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <-sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//当前连接的业务
	// fmt.Println("连接建立成功")

	user := NewUser(conn)

	//用户上线，将用户加入到onlineMap中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	
	//广播当前用户上线消息
	this.BroadCast(user,"已上线")

	//当前handler阻塞
	select {}

}

//启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d",this.Ip, this.Port))
	if err != nil {
		fmt.Println("监听状态启动失败",err)
		return
	}

	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListenMessage()

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