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
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(val)

}

func AddOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	inte, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("wrong id for comment")
	}
	fmt.Println(r.Body)
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		panic(err)
	}
	log.Println(data)
	err2 := CommentsRestImpl.Controller.AddComment(int64(inte), data["comment"].(string), data["user"].(string))
	if err2 != nil {
		panic(err2)
	}
	log.Println("Writing comment successful")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	return_data, _ := json.Marshal("{\"message\":\"success_created\"}")
	w.Write(return_data)
}

func ModifyOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	inte, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("wrong id for comment")
	}
	fmt.Println(r.Body)
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		panic(err)
	}
	log.Println(data)
	err2 := CommentsRestImpl.Controller.ModifyComment(int64(inte), data["comment"].(string))
	if err2 != nil {
		panic(err2)
	}
	log.Println("Modifying comment successful")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	return_data, _ := json.Marshal("{\"message\":\"success_changes\"}")
	w.Write(return_data)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	inte, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("wrong id for comment")
	}
	fmt.Println(r.Body)
	err2 := CommentsRestImpl.Controller.DeleteComments(int64(inte))
	if err2 != nil {
		panic(err2)
	}
	log.Println("Deletinh comment successful")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	return_data, _ := json.Marshal("{\"message\":\"success_deletes\"}")
	w.Write(return_data)
}

func main() {
	CommentsRestImpl.Controller.AddComment(1, "comment1", "me")
	router := mux.NewRouter()
	router.HandleFunc("/comments", GetAll).Methods("GET")
	router.HandleFunc("/comment/{id}", GetOne).Methods("GET")
	router.HandleFunc("/comment/{id}", AddOne).Methods("POST")
	router.HandleFunc("/comment/{id}", ModifyOne).Methods("PUT")
	router.HandleFunc("/comment/{id}", DeleteOne).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
