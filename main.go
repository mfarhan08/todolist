package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"./model"
	"github.com/gorilla/mux"
)

func getAllTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got it")
	todos := model.GetAllTodo()
	js, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got it")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	todo := model.GetTodo(id)
	js, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got it")
	type create_request struct {
		Name string
	}
	request := create_request{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &request)
	js, _ := json.Marshal(model.CreateTodo(request.Name))
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got it")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	model.DeleteTodo(id)
	w.Write([]byte(`{"status": "Todo marked as done and deleted"}`))
}

func main() {
	model.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/todo/", getAllTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todo/", createTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", deleteTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
