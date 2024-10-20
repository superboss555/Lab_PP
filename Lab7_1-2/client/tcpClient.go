package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Подключение к TCP-серверу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close() // Закрыть соединение при завершении работы клиента

	// Чтение сообщения с клавиатуры и отправка на сервер
	fmt.Print("Enter message to send: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message := scanner.Text()

		// Отправка сообщения на сервер
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// Получение ответа от сервера
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		fmt.Print("Server response: " + response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	fmt.Println("Client is shutting down...")
}
