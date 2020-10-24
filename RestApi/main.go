package main

import(
	"fmt"
	"belajar/RestApi/controller"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"

)



func main() {
	restApi()
}

func inqueryAccounts(res http.ResponseWriter, req *http.Request){
	var account controller.Accounts
	result := account.FindAll()
	json.NewEncoder(res).Encode(result)
}

func storeAccount(res http.ResponseWriter, req *http.Request){
	var account controller.Account
	json.NewDecoder(req.Body).Decode(&account)
	// controller.Store(account)
	result := account.Create()
	// controller.Save(account)
	json.NewEncoder(res).Encode(result)
}

//update
func updateAccount(res http.ResponseWriter, req *http.Request){
	var account controller.Account
	json.NewDecoder(req.Body).Decode(&account)
	// controller.Store(account)
	result := account.Update()
	// controller.Save(account)
	json.NewEncoder(res).Encode(result)
}

//delete
func deleteAccount(res http.ResponseWriter, req *http.Request){
	var account controller.Account
	json.NewDecoder(req.Body).Decode(&account)
	// controller.Store(account)
	result := account.Delete()
	// controller.Save(account)
	json.NewEncoder(res).Encode(result)
}

//card
func inqueryCards(res http.ResponseWriter, req *http.Request){
	var cards controller.Cards
	result := cards.FindAll()
	json.NewEncoder(res).Encode(result)
}

func storeCard(res http.ResponseWriter, req *http.Request){
	var card controller.Card
	json.NewDecoder(req.Body).Decode(&card)
	// controller.Store(Card)
	result := card.Create()
	// controller.Save(Card)
	json.NewEncoder(res).Encode(result)
}

//update
func updateCard(res http.ResponseWriter, req *http.Request){
	var card controller.Card
	json.NewDecoder(req.Body).Decode(&card)
	// controller.Store(Card)
	result := card.Update()
	// controller.Save(Card)
	json.NewEncoder(res).Encode(result)
}

//delete
func deleteCard(res http.ResponseWriter, req *http.Request){
	var card controller.Card
	json.NewDecoder(req.Body).Decode(&card)
	// controller.Store(Card)
	result := card.Delete()
	// controller.Save(Card)
	json.NewEncoder(res).Encode(result)
}

func restApi(){
	fmt.Println("Server Running")
	router := mux.NewRouter().StrictSlash(true)

	//api url
	router.HandleFunc("/accounts", inqueryAccounts).Methods("GET")
	router.HandleFunc("/account", storeAccount).Methods("POST")
	router.HandleFunc("/account", updateAccount).Methods("PUT")
	router.HandleFunc("/account", deleteAccount).Methods("DELETE")
	//card
	router.HandleFunc("/cards", inqueryCards).Methods("GET")
	router.HandleFunc("/card", storeCard).Methods("POST")
	router.HandleFunc("/card", updateCard).Methods("PUT")
	router.HandleFunc("/card", deleteCard).Methods("DELETE")

	//server
	log.Fatal(http.ListenAndServe(":8080", router))
}
