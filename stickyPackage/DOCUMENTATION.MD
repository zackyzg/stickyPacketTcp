# stickyPackage

Demonstration of TCP packet sticking solution

`tcp粘包解决方案演示`

自定义编码协议，实现消息体传输可以识别一条完整数据

# 粘包解决办法

出现"粘包"的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。

封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入"包尾"内容)
。包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。

运行结果如下所示
收到的消息为完整数据，未出现的消息粘包

收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO
收到client发送的数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLO