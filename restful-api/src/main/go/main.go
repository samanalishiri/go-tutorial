package main

import (
	"fmt"
	"net/http"
	"os"
	"restfull-api/src/main/go/endpoint"
)

func main() {
	http.HandleFunc("/users", endpoint.UserEndpoint)
	http.HandleFunc("/users/", endpoint.SubUserEndpoint)
	http.HandleFunc("/", endpoint.RootEndpoint)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
