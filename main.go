package main

import (
	"fmt"

	"github.com/jedi4z/rivendell/database"
	"github.com/jedi4z/rivendell/router"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}

	r := router.Create(db)
	r.Run()
}
