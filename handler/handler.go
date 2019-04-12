package handler

import (
	"myFirstapp/controllers"

	"github.com/gorilla/mux"
)

func CreateClientHandlers(router *mux.Router) {
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
}

func OfferHandlers(router *mux.Router) {
	router.HandleFunc("/api/offers", controllers.CreateNewOffer).Methods("POST")
	router.HandleFunc("/api/offers", controllers.GetOffers).Methods("GET")
	router.HandleFunc("/api/sold_offers", controllers.SoldOffers).Methods("GET")
}

func BidHandlers(router *mux.Router) {
	router.HandleFunc("/api/bid/create", controllers.CreateBid).Methods("POST")
	router.HandleFunc("/api/bid/accept/{id}", controllers.AcceptBid).Methods("PUT")
}
