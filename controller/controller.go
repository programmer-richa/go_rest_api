package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/programmer-richa/go_rest_api/models"
	"github.com/programmer-richa/go_rest_api/repository"
	"github.com/programmer-richa/go_rest_api/utils"
)

// Controller handles all the http requests.
type Controller struct{}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var jwt models.JWT
	var error models.Error
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if user.Email == "" {
		error.Message = "Email is missing."
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}
	userRepo := repository.UserRepository{}
	user = userRepo.Login(user)
	if user == nil {
		error.Message = "The user does not exist"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}
	token, err := utils.GenerateToken(user)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwt.Token = token

	utils.ResponseJSON(w, jwt)
}

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	var error models.Error
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if user.Email == "" {
		error.Message = "Email is missing."
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}
	userRepo := repository.UserRepository{}
	success := userRepo.Signup(user)
	if !success {
		error.Message = "Account already exists"
		utils.RespondWithError(w, http.StatusBadRequest, error)
		return
	}
	error.Message = "Account Registered successfully"
	utils.ResponseJSON(w, error)

}

// Dashboard controller is middleware protected.
// This handler executes only when a request with valid token comes.
func (c *Controller) Dashboard(w http.ResponseWriter, r *http.Request) {
	var error models.Error
	error.Message = "Success."
	utils.ResponseJSON(w, error)
}
