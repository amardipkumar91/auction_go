package bid

import (
	"myFirstapp/client"
	utils "myFirstapp/dbutils"
	"myFirstapp/offer"
	u "myFirstapp/utility"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Bid struct {
	gorm.Model
	BidPrice   float32 `json:"bid_price"`
	Client     client.Client
	ClientId   uint `gorm:"ForeignKey:id"`
	Offer      offer.Offer
	OfferId    int `gorm:"ForeignKey:id" json:"offer_id"`
	IsAccepted int `json:"is_accepted" gorm:"default:'0'"`
}

func (bid *Bid) Create() map[string]interface{} {
	db := utils.GetDB()
	// if resp, ok := offer.Validate(); !ok {
	// 	return resp
	// }
	db.Create(bid)

	resp := u.Message(true, "success")
	// resp["contact"] = offer
	return resp
}

func AcceptBid(bid_id string) error {
	db := utils.GetDB()
	err := db.Exec("update bids set is_accepted = 1 where id = ?", bid_id).Error
	return err
}

func GetOfferIdFromBid(bid_id string) (offer_id int) {
	db := utils.GetDB()
	var bid Bid
	db.First(&bid, bid_id)
	return bid.OfferId
}
