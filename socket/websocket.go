package socket

import (
	"fmt"
	"log"
	"myFirstapp/offer"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *offer.Offer)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// type longLatStruct struct {
// 	Long float64 `json:"longitude"`
// 	Lat  float64 `json:"latitude"`
// }

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func Writer(coord *offer.Offer) {
	broadcast <- coord
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// register client
	clients[ws] = true
}

func Echo() {
	for {
		val := <-broadcast
		latlong := fmt.Sprintf("%d, %s, %f ", val.ID, val.Title, val.BidPrice)
		// send to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(latlong))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// func LongLatHandler(w http.ResponseWriter, r *http.Request) {
// 	var coordinates longLatStruct
// 	if err := json.NewDecoder(r.Body).Decode(&coordinates); err != nil {
// 		log.Printf("ERROR: %s", err)
// 		http.Error(w, "Bad request", http.StatusTeapot)
// 		return
// 	}
// 	defer r.Body.Close()
// 	go Writer(&coordinates)
// }
