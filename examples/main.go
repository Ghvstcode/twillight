package main

import (
	"fmt"
	"github.com/GhvstCode/twillight"
	"os"
)

func main() {
	c := twillight.NewClient("ACd4859955d9ff9fb86b0a6daabd2bd699", "f7c9a17f472979d2841c0d7a5e7495c6")
	newSms(c)
	r, err := newVerify(c)
	if err != nil {
		fmt.Print("An Error Occurred: ", err)
		os.Exit(1)
	}

	r, err = startVerification(c,r,"", "")
	if err != nil {
		fmt.Print("An Error Occurred: ", err)
		os.Exit(1)
	}
}

