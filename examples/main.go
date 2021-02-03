package main

import (
	"fmt"
	"github.com/GhvstCode/twillight"
	"os"
)


func main() {
	a := twillight.NewAuth("ACxxxxxxx", "f7xxxxxxxxx")


	l := a.NewLookupClient()

	res, err := twillight.LookupPhoneNumber(l, "+23470377777777")

	if err != nil {
		fmt.Print("An Error Occurred: ", err)
		os.Exit(1)
	}

	fmt.Println(res.CallerName)
}

