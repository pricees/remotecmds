package server

import (
	"fmt"
	"net"
)

type CmdServer struct {
	port     int
	listener net.Listener
}

func (cmdServer *CmdServer) Listen() {
	port := fmt.Sprintf(":%d", cmdServer.port)
	listener, err := net.Listen("tcp", port)
	cmdServer.listener = listener
	if err != nil {
		//handle error
	}

	go func() {
		for {
			conn, err := cmdServer.listener.Accept()
			if err != nil {
				//handle error
				continue
			}
			go cmdServer.handleConnection(conn)
		}
	}()
}

func (cmdServer *CmdServer) Close() string {
	if cmdServer.listener == nil {
		return fmt.Sprintf("Listener is not defined")
	} else {
		cmdServer.listener.Close()
		return fmt.Sprintf("Listener closed.")
	}
}

func NewCmdServer() *CmdServer {
	return &CmdServer{port: 8080}
}

func (cmdServer *CmdServer) handleConnection(conn net.Conn) {
	conn.Write([]byte("Hello World\n"))
}
