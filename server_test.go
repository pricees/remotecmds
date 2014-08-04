package server

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"testing"
	"time"
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
	exp := "Hello World\n"

	_, err := conn.Write([]byte(exp))
	if err != nil {
		//handle error
	}

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if res != exp {
		t.Errorf("Status should be %v got %v", exp, res)
	}

	server.Close()
}

func TestUTCTime(t *testing.T) {
	server := NewCmdServer()
	server.Listen()

	conn, _ := net.Dial("tcp", ":8080")

	_, err := conn.Write([]byte("UTC\n"))
	if err != nil {
		//handle error
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	validTime := regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !validTime.Match([]byte(res)) {
		t.Errorf("Time should math %s but got %v", regex, res)
	}

	server.Close()
}

/*
func TestSomething (t *testing.T) {
	server := NewCmdServer()
	server.Listen()

	conn, _ := net.Dial("tcp", ":8080")
	exp := "Hello World\n"

	_, err := conn.Write([]byte(exp))
	if err != nil {
		//handle error
	}

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if res != exp {
		t.Errorf("Status should be %v got %v", exp, res)
	}

	server.Close()
}*/
