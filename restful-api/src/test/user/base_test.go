package user

import "testing"

func checkError(t *testing.T, err error, text string) {
	if err != nil {
		t.Fatalf("%s: %s", text, err)
	}
}
