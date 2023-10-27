/**
 *
 * @package       main
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// process 处理(读取)client传过的数据
func process(conn net.Conn) {
	defer conn.Close()

	// 起一个读取器
	reader := bufio.NewReader(conn)
	var temp [1024]byte

	for {
		n, err := reader.Read(temp[:]) // 将数据写入temp中，并返回读入的字节数

		if err == io.EOF { // 读完了或者文件关闭了
			break
		}

		if err != nil {
			fmt.Println("read from client data failed,err:", err)
			break
		}

		restr := string(temp[:n])

		fmt.Println("收到client发送的数据:", restr)

	}

}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:25000") // 监听指定IP的端口

	if err != nil {
		fmt.Println("tcp Listen '127.0.0.1:25000' failed,err:", err)
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept() // 等待listen返回的conn
		if err != nil {
			fmt.Println("Accept failed,err:", err)
			continue
		}

		go process(conn)

	}

}
