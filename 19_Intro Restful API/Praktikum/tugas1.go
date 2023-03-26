package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Buat permintaan HTTP GET ke API
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		fmt.Println("Gagal melakukan permintaan ke API:", err)
		return
	}

	// Pastikan respons ditutup setelah selesai diproses
	defer resp.Body.Close()

	// Baca isi dari respons
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Gagal membaca respons dari API:", err)
		return
	}

	// Tampilkan isi dari respons
	fmt.Println(string(body))
}
