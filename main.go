package main

import (
	"fmt"

	"github.com/jedi4z/rivendell/config"
)

func main() {
	db, err := config.CreateDatabase()
	if err != nil {
		fmt.Println(err)
	}

	r := config.CreateEngine(db)
	r.Run()
}
