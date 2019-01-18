// ChatClent
package main

import (
	"fmt"
	"net"
	"os"
)

var ch chan int = make(chan int)
var nackname string

func reader(conn *net.TCPConn) {
	buff := make([]byte, 256)
	for {
		j, err := conn.Read(buff)
		if err != nil {
			ch <- 1
			break
		}
		fmt.Println("%s\n", buff[0:j])
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	TcpAdd, _ := net.ResolveTCPAddr("tcp", service)
	conn, err := net.DialTCP("tcp", nil, TcpAdd)
	if err != nil {
		fmt.Println("服务器没有开启")
		os.Exit(1)
	}
	defer conn.Close()
	go reader(conn)
	fmt.Println("请输入昵称")
	fmt.Scanln(&nackname)
	fmt.Println("你的昵称是:", nackname)

	for {
		var msg string
		fmt.Scan(&msg)
		fmt.Print("<" + nackname + ">" + "说:")
		fmt.Println(msg)
		b := []byte("<" + nackname + ">" + "说:" + msg)
		conn.Write(b)
		select {
		case <-ch:
			fmt.Println("server发生错误，请重新连接")
			os.Exit(2)
		default:
		}
	}
}
