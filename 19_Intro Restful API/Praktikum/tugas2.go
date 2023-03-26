package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	resp, err := http.Get("http://example.com/api/data")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var data []Data
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, d := range data {
		if d.ID == 3 {
			fmt.Println(d)
			break
		}
	}
}
