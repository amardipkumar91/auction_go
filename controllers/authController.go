package controllers

import (
	"encoding/json"
	"myFirstapp/client"
	u "myFirstapp/utility"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	client := &client.Client{}
	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := client.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	client := &client.Client{}
	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := client.Login(client.Email, client.Password)
	u.Respond(w, resp)
}
