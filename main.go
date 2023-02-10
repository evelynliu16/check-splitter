package main

import (
	"check-split/parse"
	"github.com/veryfi/veryfi-go/veryfi"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/uploadReceipt", parse.ParseReceipt)
	http.HandleFunc("/participants", parse.HandleParticipants)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	// Set up verifi
	parse.VeryfiServer, err = veryfi.NewClientV8(&veryfi.Options{
		ClientID: "YOUR_CLIENT_ID",
		Username: "YOUR_USERNAME",
		APIKey:   "YOUR_API_KEY",
	})
	if err != nil {
		log.Fatal(err)
	}
}
