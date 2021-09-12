package entity

import (
	"errors"
	"fmt"
)

type bank struct {
	Id                 int     `json:"id"`
	Name               string  `json:"name"`
	InterestRate       float32 `json:"interestRate"`
	MaximumLoan        float32 `json:"maximumLoan"`
	MinimumDownPayment float32
	LoanTermInMonths   uint
}

var bankList = []bank{
	{Id: 1, Name: "Oshadbank", InterestRate: 66.65, MaximumLoan: 159750, MinimumDownPayment: 20, LoanTermInMonths: 10},
	{Id: 2, Name: "Privat-bank", InterestRate: 89.96, MaximumLoan: 255620, MinimumDownPayment: 15, LoanTermInMonths: 15},
}

func GetAllBanks() []bank {
	return bankList
}

func GetBankByID(id int) (*bank, error) {
	for _, a := range bankList {
		fmt.Println("asdas")
		if a.Id == id {
			return &a, nil
		}
	}
	return nil, errors.New("Bank not found")
}
