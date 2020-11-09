package main

import (
	"fmt"
	"net/http"
	"os"
	"restfull-api/src/main/go/handler"
)


func main() {
	http.HandleFunc("/users", handler.UserRouter)
	http.HandleFunc("/users/", handler.UserRouter)
	http.HandleFunc("/", handler.RootHandler)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
