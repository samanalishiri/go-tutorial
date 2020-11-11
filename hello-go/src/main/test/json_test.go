package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}
type Response struct {
	Message Message `json:"message"`
}

func Test1_Marshal(t *testing.T) {
	m := Message{"Alice", "Hello", 1294706395881547000}
	bytes, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("marshaling was failed: %v", err)
	}

	expect := []byte(`{"name":"Alice","body":"Hello","time":1294706395881547000}`)
	assert.Equal(t, len(expect), len(bytes))
	assert.Equal(t, expect, bytes)
}

func Test2_Unmarshal(t *testing.T) {
	bytes := []byte(`{"message":{"name":"Alice","body":"Hello","time":1294706395881547000}}`)
	var res Response

	err := json.Unmarshal(bytes, &res)
	if err != nil {
		t.Fatalf("unmarshaling was failed: %v", err)
	}

	assert.Equal(t, "Alice", res.Message.Name)
	assert.Equal(t, "Hello", res.Message.Body)
	assert.Equal(t, int64(1294706395881547000), res.Message.Time)
}
