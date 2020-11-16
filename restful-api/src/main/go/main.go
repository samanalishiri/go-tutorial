package main

import (
	"fmt"
	"net/http"
	"os"
	"restfull-api/src/main/go/root"
	"restfull-api/src/main/go/user"
)

var dispatcher = user.NewDispatcher()

func main() {
	dispatcher := dispatcher.Init()
	http.HandleFunc("/users", dispatcher.Dispatcher)
	http.HandleFunc("/users/", dispatcher.Dispatcher)
	http.HandleFunc("/", root.RootEndpoint)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
