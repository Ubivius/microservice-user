package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-user/pkg/data"
	"github.com/gorilla/mux"
)

func TestValidationMiddlewareWithValidBody(t *testing.T) {
	// Creating request body
	body := &data.User{
		Username:    "addName",
		Email:       "user@email.com",
		DateOfBirth: "01/01/1999",
	}
	bodyBytes, _ := json.Marshal(body)

	request := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	userHandler := NewUsersHandler(newUserDB())

	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.AddUser)
	router.Use(userHandler.MiddlewareUserValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestValidationMiddlewareWithNoUsername(t *testing.T) {
	// Creating request body
	body := &data.User{
		Email:       "user@email.com",
		DateOfBirth: "01/01/1999",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Error("Body passing to test is not a valid json struct : ", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	userHandler := NewUsersHandler(newUserDB())

	// Create a router for middleware because linking is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.AddUser)
	router.Use(userHandler.MiddlewareUserValidation)

	// Server http on our router
	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Field validation for 'Username' failed on the 'required' tag") {
		t.Error("Expected error on field validation for Name but got : ", response.Body.String())
	}
}
