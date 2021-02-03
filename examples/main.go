package main

import (
	"fmt"
	"github.com/GhvstCode/twillight"
	"os"
)


func main() {
	a := twillight.NewAuth("ACd4859955d9ff9fb86b0a6daabd2bd699", "f7c9a17f472979d2841c0d7a5e7495c6")

	l := a.NewLookupClient()

	res, err := twillight.LookupPhoneNumber(l, "+23470377777777")

	if err != nil {
		fmt.Print("An Error Occurred: ", err)
		os.Exit(1)
	}

	fmt.Println(res.CallerName)
}

