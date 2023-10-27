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

func process(conn net.Conn) {
	defer conn.Close()

	// 起一个读取器，读取收到的数据
	var temp [1024]byte
	reader := bufio.NewReader(conn)
	for {
		n, err := reader.Read(temp[:]) // Read将数据读入temp。它返回读入temp的字节数

		if err == io.EOF { // io.EOF（表示已经到达文件的末尾或连接已关闭）
			break
		}

		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}

		nowstr := string(temp[:n])
		fmt.Println("收到client数据:", nowstr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen tcp `127.0.0.1:20000` failed,err:", err)
		return
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed,err:", err)
			continue
		}

		go process(conn)

	}

}
