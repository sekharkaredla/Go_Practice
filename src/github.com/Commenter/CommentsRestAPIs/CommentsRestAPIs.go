package main

import (
	json "encoding/json"
	"fmt"
	log "log"
	http "net/http"
	strconv "strconv"

	CommentsRestImpl "github.com/Commenter/CommentsRest"
	mux "github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CommentsRestImpl.Controller.GetAllComments())
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	inte, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("wrong id for comment")
	}
	err, val := CommentsRestImpl.Controller.GetComment(int64(inte))
	if err != nil {
		log.Fatal("Wrong id")
	}
	fmt.Println(json.Marshal(val))
	json.NewEncoder(w).Encode(val)
}

func AddOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	_, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("wrong id for comment")
	}
	type Temp struct {
		id      int    `json:"id"`
		content string `json:"content"`
		user    string `json:"user"`
	}
	var temp1 Temp
	json.NewDecoder(r.Body).Decode(&temp1)
	fmt.Println(temp1)
}

func main() {
	CommentsRestImpl.Controller.AddComment(1, "comment1", "me")
	router := mux.NewRouter()
	router.HandleFunc("/comments", GetAll).Methods("GET")
	router.HandleFunc("/comment/{id}", GetOne).Methods("GET")
	router.HandleFunc("/comment/{id}", AddOne).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
