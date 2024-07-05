package cmd

import (
	"log"
	"net"
	"time"
)

func WaitForService(host string) {
	log.Printf("waiting for %s", host)

	for {
		log.Printf("testing connection for %s", host)
		con, err := net.Dial("tcp", host)
		if err == nil {
			_ = con.Close()
			log.Printf("%s is running!", host)
			return
		}
		time.Sleep(time.Millisecond * 500)
	}
}
