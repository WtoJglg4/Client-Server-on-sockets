package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	logFile, err := os.OpenFile("server_log.txt", os.O_WRONLY, os.FileMode(os.O_APPEND)) //открываем log файл
	if err != nil {
		panic(err)
	}
	defer logFile.Close() //отложенное закрытие

	errorLog := log.New(logFile, "ERR: ", log.Ltime) //логгирование ошибок

	ln, err := net.Listen("tcp", ":2020") //слушаем 2020 порт
	if err != nil {
		errorLog.Panic(err)
	}
	logFile.WriteString("INFO: server started " + time.Now().String() + "\n") //логгируем "сервер запущен"

	count := 0 //счетчик подключеных клиентов
	for {
		conn, err := ln.Accept() //создали подключение
		if err != nil {
			errorLog.Panic(err)
		}
		logFile.WriteString("INFO: client connected " + time.Now().String() + "\n") //логгируем "подключился клиент"
		go handleClient(conn, &count, logFile)                                      //(горутина)работаем сразу с несколькими клиентами(передали объект conn, указатель на счетчик, указатель на лог файл)
	}
}

func handleClient(conn net.Conn, count *int, logFile *os.File) {
	defer conn.Close()                                                                             //отложенное закрытие подключения
	defer logFile.WriteString("INFO: client has been diconnected " + time.Now().String() + "\n\n") //логгируем "клиент отключен"

	data := make([]byte, 100) //создаем дин.массив байт
	for i := 0; i < len(data); i++ {
		data[i] = 32 //инициализируем пробелами(32 в utf-8 = пробел)
	}
	conn.Read(data) //читаем из подключения данные
	*count++        //прибавляем счетчик клиентов

	//"укорачиваем" массив данных, избавляясь от лишних пробелов в конце массива
	index := len(data) - 1
	for data[index] == 32 {
		data = data[:index]
		index--
	}

	fmt.Printf("Client_%v: %v\n", *count, string(data))                                            //вывод в консоль номер клиента и того, что он нам отправил
	logFile.WriteString("INFO: massage from client: " + string(data) + time.Now().String() + "\n") //логгируем месседж клиента

	time.Sleep(3 * time.Second)                                                                  //эмуляция работы
	Reverse(&data)                                                                               //переворачиваем строку
	data = []byte(string(data) + ". Server written by Glazov Vadim M3O-109B-22\n")               //добавляем необходимое
	conn.Write(data)                                                                             //отправляем данные клиенту
	logFile.WriteString("INFO: massage to client: " + string(data) + time.Now().String() + "\n") //логгируем то, что отправили

	time.Sleep(3 * time.Second) //эмуляция работы
}

// переворот строки
func Reverse(data *[]byte) {
	runeData := []rune(string(*data))
	k := len(runeData) / 2
	for i := 0; i < k; i++ {
		runeData[i], runeData[len(runeData)-1-i] = runeData[len(runeData)-i-1], runeData[i]
	}
	*data = []byte(string(runeData))
}
