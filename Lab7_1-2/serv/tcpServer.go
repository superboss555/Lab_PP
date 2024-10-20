package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("Received: %s\n", msg)

		// Отправляем подтверждение клиенту
		_, err := conn.Write([]byte("Message received: " + msg + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection:", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on port 8080...")

	// Прием нового соединения
	conn, err := listener.Accept() // Ожидание нового соединения
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		os.Exit(1) // Завершаем программу в случае ошибки
	}

	// Обработка нового соединения
	handleConnection(conn)

	// Закрываем сервер после обработки одного соединения
	fmt.Println("Server shutting down...")
}
