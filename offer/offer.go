package offer

import (
	"fmt"
	"myFirstapp/client"
	utils "myFirstapp/dbutils"
	u "myFirstapp/utility"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Offer struct {
	gorm.Model
	// Id       int `gorm:"primary_key"`
	BidPrice float32    `json:"bid_price"`
	GoLive   *time.Time `json:"go_live"`
	LifeTime int        `json:"life_time"`
	PhotoURL string     `json:"photo_url"`
	Title    string     `json:"title"`
	Client   client.Client
	CreateBy uint `gorm:"ForeignKey:id"`
	Sold     bool `json:"sold_offer" gorm:"default:'0'"`
}

type OfferData struct {
	Id        int
	Title     string
	CreatedAt *time.Time
	GoLive    *time.Time
	LifeTime  int
	PhotoURL  string
	BidPrice  float32
	CreatedBy string
}

func (offer *Offer) Validate() (map[string]interface{}, bool) {

	if offer.BidPrice == 0 {
		return u.Message(false, "Bid Price should be on the payload"), false
	}

	if offer.LifeTime == 0 {
		return u.Message(false, "Life Time should be on the payload"), false
	}

	if offer.PhotoURL == "" {
		return u.Message(false, "Photo URL should be on the payload"), false
	}

	if offer.Title == "" {
		return u.Message(false, "Title should be on the payload"), false
	}
	return u.Message(true, "success"), true
}

func (offer *Offer) Create() map[string]interface{} {
	db := utils.GetDB()
	if resp, ok := offer.Validate(); !ok {
		return resp
	}
	db.Create(offer)

	resp := u.Message(true, "success")
	// resp["contact"] = offer
	return resp
}

func (offer *Offer) Query(page int, size int, sortkey string, is_sold int) ([]*OfferData, error) {
	fmt.Println("start the query from the ")

	if size == 0 {
		size = 10
	}
	var offsetVal int = 0
	if page > 0 {
		offsetVal = (page * size)
	} else {
		offsetVal = 0
	}

	if sortkey == "" {
		sortkey = "go_live"
	}
	fmt.Println(offsetVal)
	// var offers []*Offer
	db := utils.GetDB()
	var result []*OfferData
	err := db.Raw("select t1.id, t1.title,t1.created_at, t1.go_live, t1.life_time, t1.photo_url, t1.bid_price, t2.email as created_by from offers as t1 join clients as t2 on t1.create_by = t2.id where sold = ? order by go_live desc limit ? offset ?", is_sold, size, offsetVal).Scan(&result).Error
	// err := db.Order(sortkey + " " + "desc").Limit(size).Offset(offsetVal).Find(&offers).Error
	return result, err
}

func GetOfferById(offer_id int) (*Offer, error) {
	db := utils.GetDB()
	var offers []*Offer

	err := db.Where("id = ?", offer_id).Find(&offers).Error
	fmt.Println("----->>>>", offers)
	if len(offers) > 0 {
		return offers[0], err
	}
	return nil, err
}

func UpdateOffer(offer_id int, bid_price float32) {
	db := utils.GetDB()
	db.Exec("update offers set bid_price = ? where id = ?", bid_price, offer_id)
}

func SoldOffer(offer_id int) error {
	db := utils.GetDB()
	err := db.Exec("update offers set sold = ? where id = ?", 1, offer_id).Error
	return err
}
