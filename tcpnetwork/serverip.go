package main 

import (
	"net"
	"log"
	"fmt"
)
func main(){
	server, err := net.Listen("113.161.65.37:3000")
	if err != nil { 
		log.Fatal(err)
	}
	for {
		_, err := server.Accept()
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(" da ket noi tcp network _ server_ client ")
	}

}
