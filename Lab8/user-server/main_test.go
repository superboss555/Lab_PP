package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func init() {
	initTestDatabase() // Инициализация тестовой базы данных
	db = testDB        // Установка глобальной переменной db на testDB
	clearTestDatabase()
}

func initTestDatabase() {
	var err error
	dsn := "host=localhost user=КотЭсНожом password=12340 dbname=userdb_test port=5432 sslmode=disable"
	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	testDB.AutoMigrate(&User{})
	clearTestDatabase()
}

func clearTestDatabase() {
	testDB.Exec("DELETE FROM users") // Очищаем таблицу перед каждым тестом
}

func TestCreateUser(t *testing.T) {
	clearTestDatabase() // Очищаем базу данных перед тестом

	newUser := User{Name: "Test User", Email: "test@example.com"}
	jsonData, _ := json.Marshal(newUser)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		return
	}

	var user User
	json.NewDecoder(rr.Body).Decode(&user)

	// Проверка, что ID должен быть равен 1
	if user.ID != 1 {
		t.Errorf("expected user ID to be 1, got %v", user.ID)
	}
}

func TestGetUsers(t *testing.T) {
	clearTestDatabase()
	req, err := http.NewRequest("GET", "/users?page=1&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUserByID(t *testing.T) {
	clearTestDatabase() // Очистка базы данных перед тестом

	// Создаем пользователя, чтобы выполнить тест
	newUser := User{Name: "Test User", Email: "test@example.com"}
	testDB.Create(&newUser) // Добавляем пользователя в базу

	// Теперь запрашиваем пользователя по ID
	req, err := http.NewRequest("GET", fmt.Sprintf("/users/%d", newUser.ID), nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUserByID) // Убедитесь, что правильно указываете функцию

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return
	}

	var user User
	if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	// Проверка, что ID пользователя соответствует ожидаемому
	if user.ID != newUser.ID {
		t.Errorf("expected user ID to be %d, got %d", newUser.ID, user.ID)
	}
}

func TestUpdateUser(t *testing.T) {
	clearTestDatabase()
	user := User{Name: "Update User", Email: "update@example.com", Age: 25}
	testDB.Create(&user)

	user.Name = "Updated User"
	body, _ := json.Marshal(user)

	req, err := http.NewRequest("PUT", "/users/"+strconv.Itoa(user.ID), bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseUser User
	json.NewDecoder(rr.Body).Decode(&responseUser)

	if responseUser.Name != user.Name {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseUser.Name, user.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	clearTestDatabase() // Очистка базы данных перед тестом

	user := User{Name: "Delete User", Email: "delete@example.com", Age: 30}
	testDB.Create(&user)

	req, err := http.NewRequest("DELETE", "/users/"+strconv.Itoa(user.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestGetUserNotFound(t *testing.T) {
	clearTestDatabase() // Очистка базы данных перед тестом

	req, err := http.NewRequest("GET", "/users/99999", nil) // Используем не существующий ID
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUserByID) // Убедитесь, что это корректный обработчик

	handler.ServeHTTP(rr, req)

	// Проверка статус-кода
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// Проверка на содержание ответа
	expectedBody := `{"error":"User not found"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %q want %q", rr.Body.String(), expectedBody)
	}
}
