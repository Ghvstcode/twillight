package main

import (
	"fmt"
	"github.com/Ghvstcode/twillight"
	"github.com/Ghvstcode/twillight/internal/verify"
)

func exampleVerifyService(auth *twillight.Auth,FriendlyName string, opts twillight.VerOptions)(*verify.ResponseVerifyService, error){
	exampleNewVerifyService, err := auth.NewVerificationService(FriendlyName, opts)
	if err != nil{
		fmt.Println(err)
	}
	return exampleNewVerifyService, nil
}

func exampleVerify(svc *verify.ResponseVerifyService, to, channel string){
	v, err := twillight.StartVerification(svc, to, channel)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(v.Sid)
}

func exampleConfirmVerify(svc *verify.ResponseVerifyService, to, code string)string{
	res, err := twillight.CompleteVerification(svc, to, code)
	if err != nil{
		fmt.Println(err)
	}

	return res.Status
}
