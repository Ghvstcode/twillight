package main

import (
	"fmt"
	"github.com/Ghvstcode/twillight"
	"github.com/Ghvstcode/twillight/internal/lookup"
)

func exampleLookupService(auth *twillight.Auth)*lookup.ClientLookup{
	return auth.NewLookupClient()
}

func examplePhoneLookup(client *lookup.ClientLookup, phone string) string{
	l, err := twillight.LookupPhoneNumber(client, phone)
	if err != nil{
		fmt.Println(err)
	}

	return l.CountryCode
}