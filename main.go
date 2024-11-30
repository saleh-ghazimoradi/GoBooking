package main

import (
	"github.com/saleh-ghazimoradi/GoBooking/config"
	"github.com/saleh-ghazimoradi/GoBooking/utils"
)

func main() {
	env := config.EnvConfig()
	_, err := utils.PostConnection(env)
	if err != nil {
		panic(err)
	}
}
