package main

import (
	"github.com/gorilla/mux"

	"net/http"

	"io/ioutil"

	"fmt"
)

func readContent() ([]byte, error) {
	data, err := ioutil.ReadFile("/data/tero.txt")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		data, err := readContent()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"data": "%s"}`, string(data))))
	})

	panic(http.ListenAndServe(":8080", r))
}
