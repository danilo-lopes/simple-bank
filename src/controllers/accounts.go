package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/danilo-lopes/simple-bank/src/db"
	"github.com/danilo-lopes/simple-bank/src/engine"
	"github.com/danilo-lopes/simple-bank/src/models"
	"github.com/danilo-lopes/simple-bank/src/repositories"
	"github.com/danilo-lopes/simple-bank/src/responses"
	"github.com/gorilla/mux"
)

var (
	defaultAccountLimit int = 10000
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro := json.Unmarshal(body, &user); erro != nil {
		responses.JSON(w, http.StatusBadRequest, erro)
		return
	}

	var account models.Account
	account.Active = true
	account.AvailableLimit = defaultAccountLimit
	user.Account = account

	if violation := engine.AuthorizeAccountCreation(user); len(violation.Violations) > 0 {
		responses.JSON(w, http.StatusUnauthorized, violation.Violations)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewUsersRepository(db)
	user.ID, erro = repository.CreateAccount(user)
	defer db.Close()
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, erro)
	}

	responses.JSON(w, http.StatusAccepted, user)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["userID"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewUsersRepository(db)
	user, erro := repository.GetAccount(userID)
	defer db.Close()
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
