package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"

	"net/http"
)

type (
	db struct {
		conn       *mgo.Session
		collection *mgo.Collection
	}

	handler struct {
		db *db
	}

	// User data
	User struct {
		ID      string `json:"id,omitempty",bson:"-"`
		Name    string `json:"name",bson:"name"`
		Surname string `json:"surname",bson:"surname"`
	}
)

func (h *handler) saveUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`Empty body`))
		return
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := h.db.collection.Insert(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (h *handler) getUsers(w http.ResponseWriter, r *http.Request) {
	users := make([]User, 0)

	err := h.db.collection.Find(bson.M{}).All(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	b, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}

func main() {
	session, err := mgo.Dial("mongodb")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	if err := session.Ping(); err != nil {
		panic(err)
	}

	log.Print("Connection to mongo success!")

	handler := &handler{
		db: &db{
			conn:       session,
			collection: session.DB("users").C("users"),
		},
	}

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.saveUser).Methods(http.MethodPost)
	r.HandleFunc("/users", handler.getUsers).Methods(http.MethodGet)

	panic(http.ListenAndServe(":8080", r))
}
