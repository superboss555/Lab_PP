package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	activeConnections sync.WaitGroup
	shutdownChannel   chan os.Signal
	mu                sync.Mutex
	connections       []net.Conn
)

func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	defer activeConnections.Done()

	mu.Lock()
	connections = append(connections, conn)
	mu.Unlock()

	fmt.Println("Client connected:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Closing connection:", conn.RemoteAddr())
			return
		default:
			if scanner.Scan() {
				msg := scanner.Text()
				fmt.Printf("Received: %s\n", msg)

				_, err := conn.Write([]byte("Message received: " + msg + "\n"))
				if err != nil {
					fmt.Println("Error sending message:", err)
					return
				}
			} else {
				fmt.Println("Client disconnected:", conn.RemoteAddr())
				return
			}
		}
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

	shutdownChannel = make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Горутина для обработки сигнала завершения
	go func() {
		<-shutdownChannel
		fmt.Println("Shutting down server...")

		// Отключение всех активных клиентов
		mu.Lock()
		for _, conn := range connections {
			conn.Close()
		}
		mu.Unlock()

		cancel()         // Отмена контекста
		listener.Close() // Закрытие слушателя завершает Accept()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return // Возврат здесь позволит выйти из цикла, если произошла ошибка
		}

		activeConnections.Add(1)
		go handleConnection(ctx, conn)

		// Проверка на завершение контекста
		if ctx.Err() != nil {
			break // Выходим из цикла, если контекст отменен (при получении сигнала завершения)
		}
	}

	// Ожидаем завершения всех активных соединений перед завершением
	activeConnections.Wait()
	fmt.Println("All connections closed. Server shutdown complete.")
}
