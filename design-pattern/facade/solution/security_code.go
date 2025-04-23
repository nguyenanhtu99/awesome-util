package main

import (
	"errors"
	"fmt"
)

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return errors.New("security code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}
