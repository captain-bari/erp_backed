package rest

import (
	log "erp/log"
	"net/http"

	"github.com/rs/cors"
)

func Init() error {

	log.Logger.Info("Starting rest server")

	mux := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodOptions,
			http.MethodConnect,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)
	mux.HandleFunc("/v1/login", login)
	mux.HandleFunc("/v1/getmaterials", getMaterials)

	handler := cors.Handler(mux)
	// err := http.ListenAndServeTLS(":8090", "/app/cert.pem", "/app/key_open.pem", handler)
	// if err != nil {
	// 	log.Logger.Errorf("WebInterface: ListenAndServe: %s", err.Error())
	// 	return err
	// }
	err := http.ListenAndServe(":8090", handler)
	if err != nil {
		log.Logger.Errorf("WebInterface: ListenAndServe: %s", err.Error())
		return err
	}
	return nil
}
