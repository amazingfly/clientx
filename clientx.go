package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func menu()  {

	input := 0

	list := []string {
		"1. Smell Control",
		"2, Lighting"}
	
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

    switch {
    case input == 1:
        MenuSmell()
	break
    case input == 2:
	MenuLight()
	break
/*    case default:
    	MenuBack()*/
	
	}
}

func MenuSmell() {
	
	input := 0

	fmt.Println("Howdy folks! What would you like to do about this smell?")
	
	list := []string {
		"1. Turn smell on!",
		"2. Turn smell off!"}

	for i := 0; i < len(list); i++ {
	fmt.Println(list[i])
	}

	switch {
	case input == 1:
		go SendCommand("outlet1on")
		break
	case input == 2:
		go SendCommand("outlet1off")
		break
	}
}

func MenuLight(){
	
}

func SendCommand(cmd string){

	conn, _ := net.Dial("tcp", "localhost:8484")
	b := make([]byte, 0)
	b =  []byte(cmd)
	fmt.Println(b)
	//for i := 0; i <4; i++{
	//		b = append(b, 1)
	//}
	fmt.Println(conn.Write(b))
	conn.Close()

}		

func main(){
	fmt.Println("Welcome to Derek's house! What would you like to control?")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Run command")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	

	b := make([]byte, 0)
	//b =  []byte()
	for i := 0; i < len(text) - 2; i++ {
		b = append(b, text[i])
	}

	conn, _ := net.Dial("tcp", "192.168.0.10:8484")
	//b := make([]byte, 0)
	//b =  []byte(text)
	fmt.Println(b)
	//for i := 0; i <4; i++{
	//		b = append(b, 1)
	//}
	fmt.Println(conn.Write(b))
	conn.Close()

}
