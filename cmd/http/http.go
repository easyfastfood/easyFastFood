//go:generate go run http.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"easyfood/config/http"
	"easyfood/pkg/graphql"
)

func main() {
	server := http.Server{
		Addr:    config.GraphqlAddr,
		Handler: graphql.NewHandler(),
	}

	fmt.Printf("Listening on %s", config.GraphqlAddr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
