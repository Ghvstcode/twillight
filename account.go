package twillight

import (
	"github.com/GhvstCode/twillight/internal/account"
)

//GetAccountInfo Returns account information
func (c *Auth) GetAccountInfo() (*account.ResponseAccount, error) {
	res, err := account.InternalGetAccountInfo(c.Client)
	return res, err
}
