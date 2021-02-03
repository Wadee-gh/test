package main

import (
	//"encoding/json"
    "log"
	"net/http"
    "github.com/gorilla/mux"
)

//create a user modle to define the user attributes

type user struct {
	ID int `json:ID`
	UserName string `json:UserName` 
	Application string `json:Application`
	Comment string `json:Comment`
	Budget int `json:Budget`
}

//empty slice of user
var users []user

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getusers).Methods("Get")  
	router.HandleFunc("/users/{id}", getuser).Methods("Get")  
	router.HandleFunc("/users", adduser).Methods("Post")  
	router.HandleFunc("/users", updateuser).Methods("Put")  
	router.HandleFunc("/users/{id}", removeuser).Methods("Delete")  

// Start the server
    http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000",router))
}

func getusers(w http.ResponseWriter, r *http.Request)  {
	log.Println("Get all users")	
}

func getuser(w http.ResponseWriter, r *http.Request)  {
	log.Println("Get a single user by id")	
}

func adduser(w http.ResponseWriter, r *http.Request)  {
	log.Println("add newuser")	
}

func updateuser(w http.ResponseWriter, r *http.Request)  {
	log.Println("Update user ")	
}

func removeuser(w http.ResponseWriter, r *http.Request)  {
	log.Println("delete user is called")
		
}

