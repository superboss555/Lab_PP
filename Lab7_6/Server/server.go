package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Устанавливаем Upgrader для upgrade HTTP-соединения до WebSocket-соединения.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Разрешаем все источники (для тестов)
	},
}

// Client представляет собой подключенного клиента
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

// ClientManager управляет подключёнными клиентами
type ClientManager struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mu         sync.Mutex
}

// NewClientManager создает новый ClientManager
func NewClientManager() *ClientManager {
	return &ClientManager{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

// Start запускает управление клиентами
func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-manager.register:
			manager.mu.Lock()
			manager.clients[client] = true
			manager.mu.Unlock()
			fmt.Println("New client connected")

		case client := <-manager.unregister:
			manager.mu.Lock()
			if _, ok := manager.clients[client]; ok {
				delete(manager.clients, client)
				close(client.send)
			}
			manager.mu.Unlock()
			fmt.Println("Client disconnected")

		case message := <-manager.broadcast:
			manager.mu.Lock()
			for client := range manager.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(manager.clients, client)
				}
			}
			manager.mu.Unlock()
		}
	}
}

// handleConnections обрабатывает входящие соединения WebSocket
func handleConnections(manager *ClientManager, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	client := &Client{conn: conn, send: make(chan []byte)}
	manager.register <- client

	go client.writeMessages()
	client.readMessages(manager)
}

// (client) readMessages читает сообщения от клиента
func (client *Client) readMessages(manager *ClientManager) {
	defer func() {
		manager.unregister <- client
		client.conn.Close()
	}()
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			fmt.Println("Error while reading message:", err)
			return
		}
		// Рассылаем полученное сообщение всем клиентам
		manager.broadcast <- message
	}
}

// (client) writeMessages пишет сообщения клиенту
func (client *Client) writeMessages() {
	for message := range client.send {
		err := client.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("Error while writing message:", err)
			return
		}
	}
}

func main() {
	manager := NewClientManager()
	go manager.Start()

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		handleConnections(manager, w, r)
	})

	fmt.Println("Chat server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
