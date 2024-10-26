package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"fmt"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User структура для представления пользователя
type User struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// DB переменная для хранения ссылки на базу данных
var db *gorm.DB
var err error

// ErrorResponse структура для возврата ошибок
type ErrorResponse struct {
	Message string `json:"message"`
}

// initDatabase инициализация базы данных
func initDatabase() {
	var err error
	dsn := "host=localhost user=КотЭсНожом password=12340 dbname=userdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	db.AutoMigrate(&User{}) // Автоматическая миграция для создания таблицы
}

// handleError централизованная обработка ошибок
func handleError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
}

// getUsers обработчик для получения списка пользователей
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	var totalUsers int64

	// Получение параметров запроса.
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	nameFilter := r.URL.Query().Get("name")

	page, limit := 1, 10 // Значения по умолчанию.

	// Парсинг параметров.
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	// Формирование запроса с учетом фильтрации.
	query := db.Model(&User{})
	if nameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+nameFilter+"%")
	}

	// Получение общего количества пользователей.
	query.Count(&totalUsers)

	// Получение пользователей с пагинацией.
	offset := (page - 1) * limit
	query.Offset(offset).Limit(limit).Find(&users)

	// Установка заголовка и кодирования ответа.
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Total int64  `json:"total"`
		Users []User `json:"users"`
		Page  int    `json:"page"`
		Limit int    `json:"limit"`
	}{
		Total: totalUsers,
		Users: users,
		Page:  page,
		Limit: limit,
	}

	json.NewEncoder(w).Encode(response)
}

// getUserByID обработчик для получения пользователя по ID
func getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User

	// Используем vars["id"], чтобы получить значение ID из маршрута
	if result := db.First(&user, vars["id"]); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Если пользователь не найден, возвращаем 404
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound) // Устанавливаем статус 404
			response := `{"error":"User not found"}`
			w.Write([]byte(response)) // Используем []byte для отправки как есть
			return
		}
		// Обработка других ошибок базы данных
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Если пользователь найден, возвращаем его данные
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// createUser обработчик для добавления нового пользователя
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	// Валидация данных
	if user.Name == "" || user.Email == "" {
		handleError(w, fmt.Errorf("Name and Email are required"), http.StatusBadRequest)
		return
	}

	db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// updateUser обработчик для обновления пользователя
func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User

	// Проверка существования пользователя
	if result := db.First(&user, vars["id"]); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			handleError(w, fmt.Errorf("Пользователь не найден"), http.StatusNotFound)
			return
		}
		handleError(w, result.Error, http.StatusInternalServerError)
		return
	}

	// Декодирование новых данных пользователя
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	// Валидация данных
	if user.Name == "" || user.Email == "" {
		handleError(w, fmt.Errorf("Имя и Email обязательны для заполнения"), http.StatusBadRequest)
		return
	}

	// Обновление пользователя
	if err := db.Save(&user).Error; err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// deleteUser обработчик для удаления пользователя
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User

	// Проверка существования пользователя
	if result := db.First(&user, vars["id"]); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			handleError(w, fmt.Errorf("Пользователь не найден"), http.StatusNotFound)
			return
		}
		handleError(w, result.Error, http.StatusInternalServerError)
		return
	}

	// Удаление пользователя
	if err := db.Delete(&user).Error; err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // Возвращаем статус 204 No Content
}

// main функция
func main() {
	initDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUserByID).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
