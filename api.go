package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct{
	listenAddress string
}

//create the server
func NEWAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}

//run the server
//Mux router HandleFunc second param must be a HTTP handler to match 
//We are going to decorate our handleAccount function to be a HTTP Handler and handle error there
//--> create a type function signature and func decorator

func (s *APIServer) Run() {
	router := mux.NewRouter()
	
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	
	log.Println("Server running on port", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)

}

//this func will derive to specific function depending on HTTP method
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error{
	if r.Method == "GET" {
		return s.handleGetAccount(w,r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w,r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w,r)
	}
	return fmt.Errorf("method not allowed %s,", r.Method)
}


func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error{
	account := NewAccount("Rick","Sanchez")
	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error{
	return nil
}

//HandlerFunc decorator impl to handle error
type apiFunc func(http.ResponseWriter, *http.Request) error 
type apiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w,r); err != nil {
			writeJSON( w, http.StatusBadRequest, apiError{ Error: err.Error()} )
		}
	}
}
//

//Func to set Header and send JSON-formatted responses
func writeJSON(w http.ResponseWriter, status int , v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)

} 

