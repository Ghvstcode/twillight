package main

import (
	"fmt"
	"github.com/Ghvstcode/twillight"
	"github.com/Ghvstcode/twillight/internal/app"
)


func main() {
	a := twillight.NewAuth("ACxxxxxxx", "f7xxxxxxxxx")

	smsClient := a.NewSmsClient()

	res, err := twillight.NewOutgoingMessage(smsClient, "+443566778", "+1543222", "HelloWorld")

	if err != nil {
		er := err.(*app.ErrorResponse)
			fmt.Println("An Error Occured! status Code is", er.ErrorCode())

		fmt.Println(err.Error())
	}

	fmt.Println(res.Status)
	//l := a.NewLookupClient()
	//
	//res, err := twillight.LookupPhoneNumber(l, "+23470377777777")
	//
	//if err != nil {
	//	fmt.Print("An Error Occurred: ", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(res.CallerName)
}

