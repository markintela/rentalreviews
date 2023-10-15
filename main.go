package main

import (
	"flag"
	"fmt"
	"log"
	api "rentalreviewspt/api/server"
	storage "rentalreviewspt/database"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("<h1>Room Rental Reviews PT </h1>"))
	// })
	// http.ListenAndServe(":8080", nil)
	listenAddr := flag.String("listenaddr", ":8080", "Server Address")
	flag.Parse()

	store := storage.NewMemoryStorage()
	server := api.NewServer(*listenAddr, store)
	fmt.Println("Server running")
	log.Fatal(server.Start())
}
