package twillight

import (
	"github.com/Ghvstcode/twillight/internal/account"
)

//GetAccountInfo Returns account information
func (a *Auth) GetAccountInfo() (*account.ResponseAccount, error) {
	res, err := account.InternalGetAccountInfo(a.Client)
	return res, err
}
