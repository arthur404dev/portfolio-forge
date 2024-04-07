package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// In-memory placeholder for resume data
var resumeData = map[string]interface{}{
	"about":       "Enthusiastic Go developer...",
	"experiences": []interface{}{},
}

// Interfaces (to be refined later)
type ResumeService interface {
	GetResumeData() (map[string]interface{}, error)
}

type MemoryResumeService struct{}

func (s *MemoryResumeService) GetResumeData() (map[string]interface{}, error) {
	return resumeData, nil
}

func getResumeHandler(resumeService ResumeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := resumeService.GetResumeData() // Ignore error for now
		json.NewEncoder(w).Encode(data)
	}
}

func main() {
	router := mux.NewRouter()
	service := &MemoryResumeService{}

	router.HandleFunc("/resume", getResumeHandler(service)).Methods("GET")

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}
