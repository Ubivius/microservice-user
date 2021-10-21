package handlers

import (
	"net/http"
)

// LivenessCheck determine when the application needs to be restarted
func (userHandler *UsersHandler) LivenessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
}

//ReadinessCheck verifies that the application is ready to accept requests
func (userHandler *UsersHandler) ReadinessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	err := userHandler.db.PingDB()
	
	if err != nil {
		log.Error(err, "DB unavailable")
		http.Error(responseWriter, "DB unavailable", http.StatusServiceUnavailable)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
