package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test1_Marshal(t *testing.T) {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)

	log.Println(err)
	log.Println(b)
}

func Test2_Unmarshal(t *testing.T) {
	b2 := []byte(`
{
    "message": {
        "name": "Alice",
        "body": "Hello",
        "time": 1294706395881547000
    }
}`)
	var m2 Response
	err2 := json.Unmarshal(b2, &m2)
	log.Println(err2)
	fmt.Printf("%+v\n", m2)
}

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}
type Response struct {
	Message Message `json:"message"`
}
