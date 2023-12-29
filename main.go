package main

import (
	"enigma-laundry/config"
	"enigma-laundry/models"
	"enigma-laundry/repository"
	"fmt"

	_ "github.com/lib/pq"
)

type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
	Address     string
}

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	con, err := config.NewDbConnection(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db := con.Conn()

	uomRepo := repository.NewUomRepository(db)
	uomRepo.Save(models.Uom{
		Id:   "1",
		Name: "Kg",
	})
}
