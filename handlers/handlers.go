package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trotelalexandre/proto/blockchain"
)

func GetBlockchain(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bc.Mux.Lock()
		defer bc.Mux.Unlock()
		json.NewEncoder(w).Encode(bc.Blocks)
	}
}

func SendTransaction(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction blockchain.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, "Invalid transaction data", http.StatusBadRequest)
			return
		}
		bc.AddTransaction(transaction)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Transaction added successfully")
	}
}

func AddWallet(bc *blockchain.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wallet := blockchain.AddWallet()
		bc.AddWallet(wallet)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(wallet)
	}
}