package main

import (
	"net"
	"fmt"
)

func main(){
	fmt.Println("Attempting to sent signal to server")

	conn, _ := net.Dial("tcp", "localhost:8484")
	b := make([]byte, 0)
	b =  []byte("outlet1on")
	fmt.Println(b)
	//for i := 0; i <4; i++{
	//		b = append(b, 1)
	//}
	fmt.Println(conn.Write(b))
	conn.Close()

}
