package server

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"testing"
	"time"
)

func setup() (*CmdServer, net.Conn) {
	server := NewCmdServer()
	server.Listen()
	conn, _ := net.Dial("tcp", ":8080")
    return server, conn
}

func TestServer(t *testing.T) {

	server := NewCmdServer()
	exp := 8080
	if server.port != exp {
		t.Errorf("Port should be %v got %v", exp, server.port)
	}
	server.Close()
}

func TestConnectionHandler(t *testing.T) {

    server, conn := setup()

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
    server, conn := setup()

	_, err := conn.Write([]byte("UTC\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	validTime := regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !validTime.Match([]byte(res)) {
		t.Errorf("Time should match %s but got %v", regex, res)
	}

	server.Close()
}

func TestCPUCurrent(t *testing.T) {
    server, conn := setup()

    cmd := "CPUCurrent"

	_, err := conn.Write([]byte(cmd+ "\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	exp:= regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !exp.Match([]byte(res)) {
		t.Errorf("Expected response to match %s but got %v", regex, res)
	}

	server.Close()
}

func TestRAM(t *testing.T) {
    server, conn := setup()

    cmd := "availableRAM"

	_, err := conn.Write([]byte(cmd + "\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	exp:= regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !exp.Match([]byte(res)) {
		t.Errorf("Expected response to match %s but got %v", regex, res)
	}

	server.Close()
}

func TestCPULastHour(t *testing.T) {
    server, conn := setup()

    cmd := "CPULastHour"

	_, err := conn.Write([]byte(cmd+ "\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	exp:= regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !exp.Match([]byte(res)) {
		t.Errorf("Expected response to match %s but got %v", regex, res)
	}

	server.Close()
}

func TestAvailableRAMLastHour(t *testing.T) {
    server, conn := setup()

    cmd := "availableRAMLastHour"

	_, err := conn.Write([]byte(cmd + "\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	exp:= regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !exp.Match([]byte(res)) {
		t.Errorf("Expected response to match %s but got %v", regex, res)
	}

	server.Close()
}

func TestDLUrl(t *testing.T) {
    server, conn := setup()

    cmd := "DownloadURL"

	_, err := conn.Write([]byte(cmd+ "\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	exp:= regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !exp.Match([]byte(res)) {
		t.Errorf("Expected response to match %s but got %v", regex, res)
	}

	server.Close()
}

func TestSay(t *testing.T) {
    server, conn := setup()

    cmd := "say"

	_, err := conn.Write([]byte(cmd+ "\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	exp:= regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !exp.Match([]byte(res)) {
		t.Errorf("Expected response to match %s but got %v", regex, res)
	}

	server.Close()
}

func TestCaptureSS(t *testing.T) {
    server, conn := setup()

    cmd := ""

	_, err := conn.Write([]byte(cmd+ "\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	exp:= regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !exp.Match([]byte(res)) {
		t.Errorf("Expected response to match %s but got %v", regex, res)
	}

	server.Close()
}

func TestTriggerWebhook(t *testing.T) {
    server, conn := setup()

    cmd := ""

	_, err := conn.Write([]byte(cmd+ "\n"))
	if err != nil {
        t.Fatal("Connection couldn't be written to.")
	}

	regex := fmt.Sprintf("%d.*UTC", time.Now().Year())
	exp:= regexp.MustCompile(regex)

	res, _ := bufio.NewReader(conn).ReadString('\n')

	if !exp.Match([]byte(res)) {
		t.Errorf("Expected response to match %s but got %v", regex, res)
	}

	server.Close()
}

