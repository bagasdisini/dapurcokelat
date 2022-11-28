package main

import (
	"app/database"
	mysql "app/pkg"
	"app/routes"
	"flag"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":3000", "http service address")

func main() {

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r)

	// fmt.Println("Running in localhost:5000")
	// http.ListenAndServe("localhost:5000", r)

	// flag.Parse()

	// room := newRoom()

	// http.HandleFunc("/", ServeHome)
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

	// 	id := r.URL.Query().Get("id")
	// 	client := NewClient(id, room, w, r)

	// 	client.room.register <- client
	// })
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
