package main

import (
	"errors"
	"fmt"
)

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return errors.New("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
