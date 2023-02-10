package parse

import (
	"fmt"
	"github.com/veryfi/veryfi-go/veryfi"
	"io"
	"net/http"
	"os"
)

type Dish struct {
	name         string
	price        float32
	participants []string
}

type Participant struct {
	name  string
	total float32
}

var VeryfiServer *veryfi.Client

func ParseReceipt(w http.ResponseWriter, r *http.Request) {

}

// Read file sent through HTTP post request
func readFileUpload(w http.ResponseWriter, r *http.Request) (error, *os.File) {
	file, header, err := r.FormFile("file")
	localFile, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err, localFile
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	}(localFile)
	if _, err = fmt.Fprintf(w, "File Name: %s\n", header.Filename); err != nil {
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err, nil
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}(localFile)
	_, err = io.Copy(localFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err, localFile
	}
	if _, err = fmt.Fprintf(w, "File uploaded successfully\n"); err != nil {
		return err, localFile
	}
	return nil, localFile
}

func HandleParticipants(w http.ResponseWriter, r *http.Request) {

}
