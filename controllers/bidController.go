package controllers

import (
	"encoding/json"
	"myFirstapp/bid"
	"myFirstapp/offer"
	u "myFirstapp/utility"
	"net/http"

	"github.com/gorilla/mux"
)

var CreateBid = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bid bid.Bid
	err := json.NewDecoder(r.Body).Decode(&bid)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	offer_id := bid.OfferId
	user_bid := bid.BidPrice
	res, err := offer.GetOfferById(offer_id)
	if res == nil {
		u.Respond(w, u.Message(false, "OfferId not found"))
		return
	}
	if res.Sold == true {
		u.Respond(w, u.Message(false, "Offer already Sold"))
		return
	}
	if user_bid <= res.BidPrice {
		u.Respond(w, u.Message(false, "User Bid Price is low from current bid price"))
		return
	} else {
		usr := r.Context().Value("user")
		bid.ClientId = usr.(uint)
		res := bid.Create()
		offer.UpdateOffer(offer_id, user_bid)
		u.Respond(w, res)
	}

}

var AcceptBid = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bid_id := params["id"]
	errAccept := bid.AcceptBid(bid_id)
	if errAccept != nil {
		u.Respond(w, u.Message(false, "Unable to Accept the Bid"))
	}
	offer_id := bid.GetOfferIdFromBid(bid_id)
	err := offer.SoldOffer(offer_id)
	if err != nil {
		u.Respond(w, u.Message(false, "Unable to Update sold in Offer"))
	}
	resp := u.Message(true, "success")
	u.Respond(w, resp)

}
