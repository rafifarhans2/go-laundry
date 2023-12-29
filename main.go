package main

import (
	"enigma-laundry/config"
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

	customer := Customer{
		Id:          "2",
		Name:        "Buda",
		PhoneNumber: "08123457",
		Address:     "Jalan rusak",
	}

	_, err = db.Exec("INSERT INTO customer VALUES ($1, $2, $3, $4)",
		customer.Id,
		customer.Name,
		customer.PhoneNumber,
		customer.Address,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("success inserting data")
}
