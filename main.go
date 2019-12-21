package main

import "github.com/gorilla/mux"

import "net/http"

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "hello"}`))
	})

	panic(http.ListenAndServe(":8080", r))
}