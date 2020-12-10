package example

import (
	"fmt"
	"net"
	"time"
)

func StartEchoClient(address string, data string) {
	//1.建立一个链接（Dial拨号）
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}

	fmt.Println("Conn Established...")

	i := 0
	bytes := make([]byte, 1024)
	for {
		// 传输数据到服务端
		send := fmt.Sprintf("data - %d", i)
		_, err = conn.Write([]byte(send))
		if err != nil {
			fmt.Printf("write failed, err:%v\n", err)
			break
		}

		_, err := conn.Read(bytes)
		if err != nil {
			fmt.Printf("read failed, err:%v\n", err)
			break
		}
		fmt.Printf("client recieve data is %s -%d\n", string(bytes), i)
		time.Sleep(1 * time.Second)
		i++
	}
}
