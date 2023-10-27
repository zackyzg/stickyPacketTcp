/**
 *
 * @package       main
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("client dial '127.0.0.1:20000' failed,err:", err)
	}

	defer conn.Close()

	// 客户端发送100次数据到服务端
	for i := 0; i < 30; i++ {
		msg := "HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease. HELLO,HELLO,HELLO"
		conn.Write([]byte(msg))
	}
}
