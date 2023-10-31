/**
 *
 * @package       stickyPackage
 * @author        YuanZhiGang <zackyuan@yeah.net>
 * @version       1.0.0
 * @copyright (c) 2013-2023, YuanZhiGang
 */

package common

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// tcpSendEncoding 发送数据编码
func TcpSendEncoding(msg string) ([]byte, error) {
	// 读取消息长度。转换成int32类型，占4个字节
	// 按照小端的方式保存消息头
	// 按照小端的方式保存消息体

	var length = int32(len(msg))
	// Buffer是一个大小可变的字节缓冲区，具有Read和Write方法。Buffer的零值是一个可供使用的空缓冲区。
	var pkg = new(bytes.Buffer)

	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

// tcpGotDecoding 接收数据解码
func TcpGotDecoding(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)           // 读取前4字节,保存的是消息大小
	now_reader := bytes.NewBuffer(lengthByte) // 起一个缓存区.
	var length int32
	err := binary.Read(now_reader, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// 检测返回数据是否符合头部信息中记录信息
	// Buffered返回可从当前缓冲区读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
