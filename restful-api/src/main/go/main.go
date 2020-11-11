package main

import (
	"fmt"
	"net/http"
	"os"
	"restfull-api/src/main/go/root"
	"restfull-api/src/main/go/user"
)

func main() {
	http.HandleFunc("/users", user.Endpoint)
	http.HandleFunc("/users/", user.FirstLevelEndpoint)
	http.HandleFunc("/", root.RootEndpoint)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
