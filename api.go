package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// implements an interface with functions that handle the interaction with db
type APIServer struct {
	listenAddress string
	store         Storage
}

// create the server
func NEWAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

//run the server
//Mux router HandleFunc second param must be a HTTP handler to match
//We are going to decorate our handleAccount function to be a HTTP Handler and handle error there
//--> create a type function signature and func decorator

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccountByID))
	router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer))


	log.Println("Server running on port", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)

}

// this func will derive to specific function depending on HTTP method on /account path
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	// if r.Method == "DELETE" {
	// 	return s.handleDeleteAccount(w, r)
	// }
	return fmt.Errorf("method not allowed %s,", r.Method)
}

func (s *APIServer) handleAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccountByID(w, r)
	}
	// if r.Method == "POST" {
	// 	return s.handleCreateAccount(w, r)
	// }
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s,", r.Method)
}

// CRUD functions
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	acc, err := s.store.GetAccount()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, acc)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	acc, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, acc)

}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	reqAccount := new(Account)
	if err := json.NewDecoder(r.Body).Decode(reqAccount); err != nil {
		return err
	}

	account := NewAccount(reqAccount.FirstName, reqAccount.SecondName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	key := []byte(os.Getenv("SECRET_KEY"))
	token, err := createJWT(account,key)
	if err != nil{
		return err
	}
	fmt.Println("token string: ",token)

	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteAccount(id); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {

		reqTransfer := new(TransferRequest)
		if err := json.NewDecoder(r.Body).Decode(reqTransfer); err != nil {
			return err
		} 
		return writeJSON(w, http.StatusOK, reqTransfer)

}

// HELPER Functions
// HandlerFunc decorator impl to handle error
type apiFunc func(http.ResponseWriter, *http.Request) error
type apiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

// Func to set Header and send JSON-formatted responses
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

//Func to get id from get request (string) and convert it to int
func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}

//Func to decorate our handlers and implement JWT authorisation
func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){




		handlerFunc(w,r)
	} 
}

//Func to check JWT 
func validateJWT(tokenString string) (*jwt.Token, error) {
	return nil, nil
}

//Func to create jwt token 
func createJWT(account *Account,secretKey []byte) (string, error) {

	// Create a new token object, specifying signing method and the desired claims
	claims := &jwt.MapClaims{
		"ExpiresAt": 15000,
		"AccountNumber": account.Number,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(secretKey)
}


