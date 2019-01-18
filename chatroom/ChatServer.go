// ChatServer
package main

import (
	"fmt"
	"net"
)

var Connmap map[string]*net.TCPConn

func checkErr(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			fmt.Println("用户退出")
			return 0
		}
		fmt.Println("发生错误")
		return -1
	}
	return 1
}

func say(tcpconn *net.TCPConn) {
	for {
		data := make([]byte, 256)
		total, err := tcpconn.Read(data)
		if err != nil {
			fmt.Println(string(data[:total]), err)
		} else {
			fmt.Println(string(data[:total]))
		}

		flag := checkErr(err)
		if flag == 0 {
			break
		}
		for _, conn := range Connmap {
			if conn.RemoteAddr().String() == tcpconn.RemoteAddr().String() {
				continue
			}
			conn.Write(data[:total])
		}
	}
}

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.5.169:8080")
	tcpListen, _ := net.ListenTCP("tcp", tcpAddr)
	Connmap = make(map[string]*net.TCPConn)
	for {
		tcpconn, _ := tcpListen.AcceptTCP()
		defer tcpconn.Close()
		Connmap[tcpconn.RemoteAddr().String()] = tcpconn
		fmt.Println("连接客户端信息:", tcpconn.RemoteAddr().String())

		go say(tcpconn)
	}
}
