package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sanitiseInput(input string) string {
	i := strings.TrimSpace(input)
	saniInput := strings.ToLower(i)
	return saniInput
}

func scannerOne(partOne chan string, saniInput string) {
	for i := 1; i < 1500; i++ {
		port := strconv.Itoa(i)
		find := net.JoinHostPort(saniInput, port)
		conn, err := net.DialTimeout("tcp", find, time.Second*3)
		handleErr(err)
		defer conn.Close()
		fmt.Fprintf(conn, "SYN")
		status, err := bufio.NewReader(conn).ReadString('\n')
		handleErr(err)
		fmt.Sprintln(find)
		if status == "ACK" {
			partOne <- find
		}
		if i == 1500 {
			partOne <- "done"
		}
	}
}

func scannerTwo(partTwo chan string, saniInput string) {
	for i := 1501; i < 3000; i++ {
		port := strconv.Itoa(i)
		find := net.JoinHostPort(saniInput, port)
		conn, err := net.DialTimeout("tcp", find, time.Second*3)
		handleErr(err)
		defer conn.Close()
		fmt.Fprintf(conn, "SYN")
		status, err := bufio.NewReader(conn).ReadString('\n')
		handleErr(err)
		fmt.Sprintln(find)
		if status == "ACK" {
			partTwo <- find
		}
		if i == 3000 {
			partTwo <- "done"
		}
	}
}

func scannerThree(partThree chan string, saniInput string) {
	for i := 3001; i < 5000; i++ {
		port := strconv.Itoa(i)
		find := net.JoinHostPort(saniInput, port)
		conn, err := net.DialTimeout("tcp", find, time.Second*3)
		handleErr(err)
		defer conn.Close()
		fmt.Fprintf(conn, "SYN")
		status, err := bufio.NewReader(conn).ReadString('\n')
		handleErr(err)
		fmt.Sprintln(find)
		if status == "ACK" {
			partThree <- find
		}
		if i == 5000 {
			partThree <- "done"
		}
	}
}

func dialTarget(saniInput string) []string {
	partOne := make(chan string, 10)
	partTwo := make(chan string, 10)
	partThree := make(chan string, 10)
	results := make([]string, 35)
	index := 0

	startingTime := time.Now()
	go scannerOne(partOne, saniInput)
	go scannerTwo(partTwo, saniInput)
	go scannerThree(partThree, saniInput)

	select {
	case res1 := <-partOne:
		results[index] = res1
		index++
		startingTime = time.Now()
	case res2 := <-partTwo:
		results[index] = res2
		index++
		startingTime = time.Now()
	case res3 := <-partThree:
		results[index] = res3
		index++
		startingTime = time.Now()
	default:
		if time.Since(startingTime) > time.Second*15 {
			panic("Timeout")
		}
	}

	return results
}

func main() {
	input := os.Args[1]
	if os.Args[1] == "" {
		fmt.Println("Requires a valid ip address")
		os.Exit(1)
	}

	saniInput := sanitiseInput(input)
	dialTarget(saniInput)
}
