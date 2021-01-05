package main

import (
	"fmt"
	"github.com/GhvstCode/twillight"
)

func newVerify(c *twillight.APIClient)(string, error){
	res, err := c.NewVerificationService("MyService")

	if err != nil {
		fmt.Print("An Error Occurred: ", err)
	}

	fmt.Println("Service created: ", res)

	return res.Sid,err
}

func startVerification(c *twillight.APIClient, serviceSid, to, channel string)(string,error){
	res, err := c.SendVerificationToken(serviceSid, "", channel)
	if err != nil {
		fmt.Print("An Error Occurred: ", err)
	}

	fmt.Println("Service created: ", res)

	return res.Sid,err
}