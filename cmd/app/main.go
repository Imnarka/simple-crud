package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

var (
	task string
	mtx  sync.Mutex
)

type TaskRequestBody struct {
	Task string `json:"task"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	defer mtx.Unlock()
	mtx.Lock()
	log.Println("Обработка GET запроса /api/hello")
	w.Write([]byte("hello, " + task))
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработка POST запроса /api/task")
	defer mtx.Unlock()
	var req TaskRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		log.Printf("Ошибка декодирования JSON: %v", err)
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}
	mtx.Lock()
	task = req.Task
	log.Printf("Переменная task изменена на %s", task)
	w.Write([]byte("Task updated"))
}

func main() {
	log.Println("Старт API сервера")
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
