package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	resp, err := http.Get("http://localhost:8083")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result Response
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(result.Message)
}
