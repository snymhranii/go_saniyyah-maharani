package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/data/", deleteDataHandler)
	http.ListenAndServe(":8080", nil)
}

func deleteDataHandler(w http.ResponseWriter, r *http.Request) {
	// Parse data ID from URL path parameter
	idStr := r.URL.Path[len("/data/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid data ID", http.StatusBadRequest)
		return
	}

	// Delete data with the given ID
	err = deleteData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data with ID %d has been deleted", id)
}

func deleteData(id int) error {
	// Perform deletion logic here
	return nil
}
