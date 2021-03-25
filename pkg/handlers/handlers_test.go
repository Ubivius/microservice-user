package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-user/pkg/data"
	"github.com/gorilla/mux"
)

// Move to util package in Sprint 9, should be a testing specific logger
func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestGetUsers(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	response := httptest.NewRecorder()

	userHandler := NewUsersHandler(NewTestLogger())
	userHandler.GetUsers(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":2") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetExistingUserByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	response := httptest.NewRecorder()

	userHandler := NewUsersHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	userHandler.GetUserByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":1") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingUserByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/users/4", nil)
	response := httptest.NewRecorder()

	userHandler := NewUsersHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	userHandler.GetUserByID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "User not found") {
		t.Error("Expected response : User not found")
	}
}

func TestDeleteNonExistantUser(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/users/4", nil)
	response := httptest.NewRecorder()

	userHandler := NewUsersHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
	}
	request = mux.SetURLVars(request, vars)

	userHandler.Delete(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "User not found") {
		t.Error("Expected response : User not found")
	}
}

func TestAddUser(t *testing.T) {
	// Creating request body
	body := &data.User{
		Username:    "player",
		Email:       "player@gmail.com",
		DateOfBirth: "01/01/1970",
	}

	request := httptest.NewRequest(http.MethodPost, "/users", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyUser{}, body)
	request = request.WithContext(ctx)

	userHandler := NewUsersHandler(NewTestLogger())
	userHandler.AddUser(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestUpdateNonExistantUser(t *testing.T) {
	// Creating request body
	body := &data.User{
		Username:    "player",
		Email:       "player@gmail.com",
		DateOfBirth: "01/01/1970",
	}

	request := httptest.NewRequest(http.MethodPut, "/users", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyUser{}, body)
	request = request.WithContext(ctx)

	userHandler := NewUsersHandler(NewTestLogger())
	userHandler.UpdateUsers(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, response.Code)
	}
}

func TestUpdateUser(t *testing.T) {
	// Creating request body
	body := &data.User{
		ID:          1,
		Username:    "player",
		Email:       "player@gmail.com",
		DateOfBirth: "01/01/1970",
	}

	request := httptest.NewRequest(http.MethodPut, "/users", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyUser{}, body)
	request = request.WithContext(ctx)

	userHandler := NewUsersHandler(NewTestLogger())
	userHandler.UpdateUsers(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingUser(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	response := httptest.NewRecorder()

	userHandler := NewUsersHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	userHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}
