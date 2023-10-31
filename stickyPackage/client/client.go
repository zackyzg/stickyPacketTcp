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
	"stickyPacketTcp/stickyPackage/common"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:25000")
	if err != nil {
		fmt.Println("dial failed,err", err)
		return
	}

	defer conn.Close()

	for i := 0; i < 21; i++ {
		msg := "HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease. HELLO,HELLO,HELLO"

		data, err := common.TcpSendEncoding(msg)
		if err != nil {
			fmt.Println("encoding msg failed,err:", err)
			return
		}

		conn.Write(data)
	}

}
