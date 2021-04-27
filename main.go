package main

import (
	"net/http"
	"DemoProject/api/userapi"

	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/find", userapi.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/get-all", userapi.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/user/create", userapi.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", userapi.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete", userapi.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/user/test", userapi.TestThu).Methods("GET")

	fmt.Printf("Golang Rest API Is Running On Port: 5000")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
}