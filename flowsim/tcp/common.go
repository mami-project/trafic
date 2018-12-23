package tcp

import (
	"fmt"
	"net"
	// "errors"
)

func closeFdSocket (conn *net.TCPConn) (error){
	f, err := conn.File()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	f.Close()
	conn.Close()
	fmt.Println("Closed fd and socket")
	return nil
}
