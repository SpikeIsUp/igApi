package router

import (
	"net/http"

	"github.com/SpikeIsUp/igApi/controller"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/aPropos", controller.APropos)
	mux.HandleFunc("/categories", controller.Categories)
	mux.HandleFunc("/collection", controller.Collection)
	mux.HandleFunc("/favoris", controller.Favoris)
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/recherche", controller.Recherche)
	mux.HandleFunc("/ressources", controller.Ressources)

	return mux
}
