package server

import (
	"fmt"
	"lab_06_01/internal/query"
	"log"
	"net/http"
)

func Run(m *query.Manager, port int) error {
	http.HandleFunc("/", home)
	http.HandleFunc("/search", search(m))

	log.Printf("Сервер запущен на http://localhost:%d", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
