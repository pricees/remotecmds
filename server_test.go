package server

import (
	"bufio"
	"net"
	"testing"
)

func TestServer(t *testing.T) {

	server := NewCmdServer()
	exp := 8080
	if server.port != exp {
		t.Errorf("Port should be %v got %v", exp, server.port)
	}
	server.Close()
}

func TestConnectionHandler(t *testing.T) {

	server := NewCmdServer()
	server.Listen()

	conn, _ := net.Dial("tcp", ":8080")
	res, _ := bufio.NewReader(conn).ReadString('\n')

	exp := "Hello World\n"

	if res != exp {
		t.Errorf("Status should be %v got %v", exp, res)
	}

	server.Close()
}
