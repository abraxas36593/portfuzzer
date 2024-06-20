package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*func defineTarget(target string)net.IP{
  t := net.ParseIP(target)
  if t == nil{
    log.Println("enter a valid ip address, url's aren't working yet.")
    os.Exit(1)
  }
  return t
}*/

func ScanTarget(okz chan string, target string) {
	writer := io.Writer(os.Stdout)
	// connecter := net.Dialer{
	for i := 1; i < 5000; i++ {
		port := strconv.Itoa(i)
		find := net.JoinHostPort(target, port)
		startTime := time.Now()
		conn, err := net.DialTimeout("tcp", find, time.Second*2)
		if time.Since(startTime) > time.Second*2 {
			continue
		}
		handleErr(err)
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, err := bufio.NewReader(conn).ReadString('\n')
		handleErr(err)
		writer.Write([]byte(find))
		writer.Write([]byte(status))
		if status == "200" {
			okz <- find
		}
	}
}

func main() {
	target := os.Args[1]
	if target == " " {
		fmt.Println("enter a valid ip address or url as a string")
		os.Exit(1)
	}
	okz := make(chan string, 100)
	go ScanTarget(okz, target)

	results := <-okz
	fmt.Println(results)
}
