package main

import (
	"enigma-laundry/config"
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

	// // SAVE
	// uomRepo.Save(models.Uom{
	// 	Id:   "3",
	// 	Name: "Pair",
	// })

	// DELETE
	if err = uomRepo.DeleteById("3"); err != nil {
		fmt.Println("failed to delete data, try again!", err)
	} else {
		fmt.Println("success delete data")
	}
}
