package controllers

import (
	"encoding/json"
	"fmt"
	"myFirstapp/offer"
	"myFirstapp/socket"
	u "myFirstapp/utility"
	"net/http"
	"strconv"
	"time"
)

var CreateNewOffer = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	usr := r.Context().Value("user")

	var offer offer.Offer
	err := json.NewDecoder(r.Body).Decode(&offer)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	fmt.Println("--", offer.GoLive)
	if offer.GoLive == nil {
		now := time.Now()
		offer.GoLive = &now
	} else {
		offer.GoLive = offer.GoLive
	}
	offer.CreateBy = usr.(uint)
	resp := offer.Create()
	go socket.Writer(&offer)
	u.Respond(w, resp)
}

var GetOffers = func(w http.ResponseWriter, r *http.Request) {
	var offer offer.Offer
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		size = 10
	}
	sortKey := r.FormValue("sort")
	fmt.Println("page and size", page, size, sortKey)
	data, err := offer.Query(page, size, sortKey, 0)
	resp := u.Message(true, "success")
	resp["offers"] = data
	defer r.Body.Close()
	// go socket.Writer(&offer)
	u.Respond(w, resp)

}

var SoldOffers = func(w http.ResponseWriter, r *http.Request) {
	var offer offer.Offer
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 0
	}
	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		size = 10
	}
	sortKey := r.FormValue("sort")
	fmt.Println("page and size", page, size, sortKey)
	data, err := offer.Query(page, size, sortKey, 1)
	resp := u.Message(true, "success")
	resp["offers"] = data
	u.Respond(w, resp)

}
