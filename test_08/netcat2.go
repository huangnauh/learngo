package main

import (
	"net"
	"log"
	"io"
	"os"
	"time"
)

func mustCopy1(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil{
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		time.Sleep(time.Second * 2)
		log.Println("done")
		done <- struct {}{}
	}()
	mustCopy1(conn, os.Stdin)
	conn.Close()
	log.Println("close")
	<-done
}
