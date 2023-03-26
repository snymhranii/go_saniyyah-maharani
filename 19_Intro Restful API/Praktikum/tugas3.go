package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	post := Post{
		Title:   "Judul Postingan",
		Content: "Isi postingan",
	}

	postJSON, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Ganti URL_API dengan URL yang diinginkan
	resp, err := http.Post("URL_API", "application/json", bytes.NewBuffer(postJSON))
	if err != nil {
		fmt.Println("Error posting data:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Postingan berhasil disimpan!")
}
