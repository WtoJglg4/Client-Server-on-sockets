package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	logFile, err := os.OpenFile("client_log.txt", os.O_WRONLY, os.FileMode(os.O_APPEND))
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	errorLog := log.New(logFile, "ERR: ", log.Ltime)

	cfgFile, err := os.Open("config.txt")
	//cfgFile, err := os.OpenFile("config.txt", os.O_RDONLY, os.FileMode(os.O_RDONLY))
	if err != nil {
		errorLog.Panic(err)
	}
	defer cfgFile.Close()

	var addr string
	fileScanner := bufio.NewScanner(cfgFile)
	for fileScanner.Scan() {
		addr += fileScanner.Text()
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		errorLog.Panic(err)
	} else {
		logFile.WriteString("INFO: client connected to server " + addr + " " + time.Now().String() + "\n")
	}

	//fmt.Printf("%T : %v", conn, conn)
	//fmt.Fprintf(conn, "HELLO, SERVER, I`M A CLIENT!\n")
	//data := []byte("Glazov Vadim M3O-109B-22")
	data := []byte("Glazov Vadim M3O-109B-22")
	time.Sleep(3 * time.Second)
	// var input string //можно добавить ввод строки для проверки асинхронности
	// fmt.Scanf("%v", &input)
	conn.Write(data)
	logFile.WriteString("INFO: message to server: " + string(data) + " " + time.Now().String() + "\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		errorLog.Panic(err)
	}
	status = string([]rune(status)[:len([]rune(status))-1]) //удаление из строки \n (для корректного логгирования)

	logFile.WriteString("INFO: message from server: " + status + " " + time.Now().String() + "\n")
	fmt.Printf("Server: %v", status)
}
