# solveMucilage

TCP packet sticking demonstration

# 为什么会出现粘包

主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。

“粘包"可发生在发送端也可发生在接收端：

由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。

# tcp粘包演示

运行结果如下所示
收到的消息出现截断，不是完整数据，出现的消息粘包

收到client数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with t
收到client数据: he disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the popul
收到client数据: ation will be infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that
the population will be infected with the disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the
population will be infected with the disease. HELLO,HELLO,HELLO
收到client数据: HELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with the disease.
HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be infected with t
收到client数据: he disease. HELLO,HELLO,HELLOHELLO,HELLO,HELLO There is a 60% probability that the population will be
infected with the disease. HELLO,HELLO,HELLO









