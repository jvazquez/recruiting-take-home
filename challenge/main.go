package main

import (
	"fmt"
	"github/jvazquez/recruiting-take-home/challenge/handlers"
	"log"
	"net/http"
	"os"
)

const defaultPort string = ":8080"

// SetupRouter will all the router with all the endpoints
func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/conversor", handlers.ConversorHandler)
	return router
}

// Api server
func main() {
	routes := SetupRouter()
	appPort := os.Getenv("PORT")
	if len(appPort) == 0 {
		appPort = defaultPort
	}
	processAddress := fmt.Sprintf(":%s", appPort)
	log.Printf("Starting to serve on %s", processAddress)
	//err := http.ListenAndServe(processAddress, routes)
	http.ListenAndServe(processAddress, routes)
	//if err != nil {
	//	log.Printf("Error is %v", err)
	//}
}
