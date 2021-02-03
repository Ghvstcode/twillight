# TWILIGHT 🦅 
Twilight is an unofficial Golang SDK for Twilio APIs. Twilight was born as a result of my inability to spell **Twilio** correctly. I searched for a **Twillio** Golang client library and couldn’t find any, I decided to build one. Halfway through building this, I realized I had spelled **Twilio** as **Twillio** when searching for a client library on Github.<br>
<br>![tenor](https://user-images.githubusercontent.com/46195831/106745073-e9a7b180-6620-11eb-99a7-cf50694f2d63.gif)


### INSTALLATION
To use this in your project, Run the following command <br>
``$ go get -u github.com/Ghvstcode/twilight`` <br>

With this project, you can interact with the following Twilio APIs <br>
 * SMS API
 * Verify API
 * Lookup API

### USAGE
The full examples can be found [here](https://github.com/Ghvstcode/twilight/blob/main/examples/main.go)<br>

* To get started with this library, you have to first authenticate. You can do this by calling the NewAuth function and passing your Account SID & Auth Token gotten from your Twilio console.<br>
* Next, You have to create a client for the API you intend to interact with. The created clients implement the interface the core functions require as their first argument<br>
* 5 digit error codes are Twilio errors. 0 error codes are my fault. You could check [here](https://www.twilio.com/docs/api/errors) for a list of all Twilio error codes<br>

This is an example showning how to use this library to send an SMS<br>

```Golang
	a := twillight.NewAuth("ACxxxxxxx", "f7xxxxxxxxx")

	smsClient := a.NewSmsClient()

	res, err := twillight.NewOutgoingMessage(smsClient, "+443566778", "+1543222", "HelloWorld")

	if err != nil {
		er := err.(*app.ErrorResponse)
			fmt.Println("An Error Occured! status Code is", er.ErrorCode())
		
		fmt.Println(err.Error())
	}
	
	fmt.Println(res.Status)
  ```

### Issues & Contributions
* You encountered any issues while using this[Or just want to mess around]? Go ahead and create a new Issue!<br>
* I would love to see this project grow, If you have any contributions or ideas, please let me know! Send in your PR's
### Notes/To-Do 
- [ ] Improve Test Coverage
- [ ] Add more API's
