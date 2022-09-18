package main

import (
	"fmt"

	"github.com/bmkgBot/internal/delivery"
)

func main() {
	fmt.Println("Apps start and running..")
	delivery.StartServer(delivery.SetupContainer())
}
