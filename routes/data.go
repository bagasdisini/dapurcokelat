package routes

import (
	"app/handlers"
	mysql "app/pkg"
	"app/repositories"

	"github.com/gorilla/mux"
)

func DataRoutes(r *mux.Router) {
	DataRepository := repositories.RepositoryData(mysql.DB)
	h := handlers.HandlerData(DataRepository)

	r.HandleFunc("/data/{dataUser}", h.ShowData).Methods("GET")
}
