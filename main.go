package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bmkgBot/internal/delivery"
)

func main() {
	delivery.StartServer(delivery.SetupContainer())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}

	fmt.Println("Apps start and running at port: " + port)
}
