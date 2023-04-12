package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sabouaram/transaction/models"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}
	defer r.Body.Close()
	if err := user.CreateUser(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	var user models.User
	user.Id = id
	if err := user.GetUser(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	respondWithJSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	user.Id = id
	if err := user.UpdateUser(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func DelUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	user := models.User{Id: id}
	if err := user.DeleteUser(); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
func MakeTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		transaction models.Transaction
		sender      models.User
		receiver    models.User
	)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request payload")
	}
	sender.Id = transaction.FromId
	receiver.Id = transaction.ToId
	if err := sender.GetUser(); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Sender ID")
	}
	if err := receiver.GetUser(); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Receiver ID")
	}
	if transaction.Amount > sender.Balance {
		respondWithError(w, http.StatusInternalServerError, "Invalid operation")
	} else {
		sender.Balance = sender.Balance - transaction.Amount
		sender.UpdateUser()
		receiver.Balance = receiver.Balance + transaction.Amount
		receiver.UpdateUser()
		respondWithJSON(w, http.StatusOK, map[string]string{"Operation": "done"})
	}
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func respondWithJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
