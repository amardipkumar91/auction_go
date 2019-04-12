package middleware

import (
	"myFirstapp/bid"
	"myFirstapp/client"
	utils "myFirstapp/dbutils"
	"myFirstapp/offer"
)

func DBAutoMigrate() {
	db := utils.GetDB()
	db.AutoMigrate(&client.Client{}, &offer.Offer{}, &bid.Bid{})
}
