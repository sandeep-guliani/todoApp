package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Todo represents a todo item
type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todoList []Todo

func initializeList() {

	todoList = append(todoList, Todo{ID: "1", Title: "Pay Credit card bill", Completed: false})
	todoList = append(todoList, Todo{ID: "2", Title: "Buy new laptop", Completed: true})

}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoList)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todoList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
		return
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Intn(1000000))
	todoList = append(todoList, todo)
	json.NewEncoder(w).Encode(&todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todoList {
		if item.ID == params["id"] {
			todoList = append(todoList[:index], todoList[index+1:]...)
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			todoList = append(todoList, todo)
			json.NewEncoder(w).Encode(&todo)
			return
		}
	}
	json.NewEncoder(w).Encode(todoList)
}

func completeTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todoList {
		if item.ID == params["id"] {
			todoList = append(todoList[:index], todoList[index+1:]...)
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			todo.Completed = true
			todoList = append(todoList, todo)
			json.NewEncoder(w).Encode(&todo)
			return
		}
	}
	json.NewEncoder(w).Encode(todoList)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todoList {
		if item.ID == params["id"] {
			todoList = append(todoList[:index], todoList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todoList)
}

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	router := mux.NewRouter()

	initializeList()

	router.HandleFunc("/todo", getTodos).Methods("GET")
	router.HandleFunc("/todo", createTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", completeTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", c.Handler(router)))
}
