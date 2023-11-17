package main

import (
	"log"
	"net/http"
	"team-work-be/config"
	_ "team-work-be/docs"
	"team-work-be/router"
	"team-work-be/utils"
	"time"
)

// @title team-work-be
// @version 1.0
// @description team-work-be

// @host localhost:18888
// @BasePath /api/v1
func main() {
	server := &http.Server{
		Addr:         "localhost:" + config.GetAppPort(),
		Handler:      router.Router(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	utils.Cronjob()

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error run server: ", err)
	}
}
