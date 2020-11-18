package main

import (
	"fmt"
	"net/http"
	"os"
	"restfull-api/src/main/go/contract"
	"restfull-api/src/main/go/root"
	"restfull-api/src/main/go/user"
)

func main() {
	mapRequest(user.NewDispatcher().Init(), "/users/", "/users")
	mapRequest(root.NewDispatcher().Init(), "/")
	start()
}

func mapRequest(dispatcher contract.FrontController, url ...string) {
	for i := 0; i < len(url); i++ {
		http.HandleFunc(url[i], dispatcher.Route)
	}
}

func start() {
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
