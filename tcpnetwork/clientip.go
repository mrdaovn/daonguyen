package main 

import (
	"net"
	"log"
	"fmt"
)
func main() {
	connection, err := net.Dial ("113.161.65.37:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(connection)

}
