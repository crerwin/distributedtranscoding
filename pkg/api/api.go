package api

import (
	"log"
	"net/http"
	"time"

	"github.com/crerwin/distributedtranscoding/pkg/dtc"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
)

// Serve brings up the DTC API
func Serve() {
	// var config config.Configuration
	port := ":" + viper.GetString("api.port")
	log.Println("Hosting API on port:", port)
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Distributed Transcoder Version " + dtc.Version))
	})
	log.Fatal(http.ListenAndServe(port, router))
}
