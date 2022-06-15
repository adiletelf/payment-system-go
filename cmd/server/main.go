package main

import (
	"github.com/adiletelf/payment-system-go/pkg/api"
	// "github.com/adiletelf/payment-system-go/pkg/models"
)


func main() {
	// db, err := api.SetupDB()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	r := api.GetRouter()
	r.Run(":8080")
}
