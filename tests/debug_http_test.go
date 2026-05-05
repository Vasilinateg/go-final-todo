package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestDebugHTTP(t *testing.T) {
	url := "http://localhost:7540/api/nextdate?now=20250701&date=20250701&repeat=y"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response: [%s]\n", string(body))
}
