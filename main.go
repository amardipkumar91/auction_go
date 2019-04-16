package main

import (
	"fmt"
	"log"
	utils "myFirstapp/dbutils"
	"myFirstapp/handler"
	"myFirstapp/middleware"
	m "myFirstapp/middleware"
	"myFirstapp/socket"
	"net/http"
	_ "net/http/pprof"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Start...")
	router := mux.NewRouter()
	utils.Init()
	m.DBAutoMigrate()
	// router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to the home page!")
	// })
	// n := negroni.Classic()
	// fmt.Println(n)
	router.HandleFunc("/", socket.RootHandler).Methods("GET")
	handler.CreateClientHandlers(router)
	handler.OfferHandlers(router)
	handler.BidHandlers(router)
	router.HandleFunc("/ws", socket.WsHandler)
	// router.HandleFunc("/longlat", socket.LongLatHandler).Methods("POST")

	router.Use(middleware.JwtAuthentication)
	go socket.Echo()

	go func() {
		log.Fatal(http.ListenAndServe(":9080", http.DefaultServeMux))
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}
	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
