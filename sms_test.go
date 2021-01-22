package twillight_test

type MockSmsService struct {
	Err error
	VerifyChannel string
	Sid string
}
