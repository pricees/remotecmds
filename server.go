/*

Go Bootcamp: Remote Commands Project

Current UTC time
Current CPU usage
Available RAM
CPU usage over last hour
Available RAM over last hour
Download url into specific folder
Make computer "say" something
Capture and send a screenshot
Trigger webhook at specific time

*/
package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
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
	input, _ := bufio.NewReader(conn).ReadString('\n')

	var res string
    cmd := strings.ToLower(input[:len(input)-1]) 
    fmt.Println(cmd)
	switch cmd {
	case "utc":
		res = time.Now().UTC().String()
    case "availableram":
        res = memoryProfile()
	default:
		res = "Hello World"
	}
	conn.Write([]byte(fmt.Sprintf("%s%c", res, '\n')))
}
func memoryProfile() string {
    return "foobar"
}
