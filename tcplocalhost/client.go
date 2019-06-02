package main 

import (
	"net"
	"log"
	"fmt"
)
func main() {
	connection, err := net.Dial ("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(connection)

}
