package entity

import (
	"errors"
)

type Bank struct {
	Id                 int     `json:"id"`
	Name               string  `form:"Name"`
	InterestRate       float64 `form:"InterestRate" `
	MaximumLoan        float64 `form:"MaximumLoan"`
	MinimumDownPayment float64 `form:"MinimumDownPayment"`
	LoanTermInMonths   uint    `form:"LoanTermInMonths"`
}

var bankList = []Bank{
	{Id: 1, Name: "Oshadbank", InterestRate: 66.65, MaximumLoan: 159750, MinimumDownPayment: 20, LoanTermInMonths: 10},
	{Id: 2, Name: "Privat-bank", InterestRate: 89.96, MaximumLoan: 255620, MinimumDownPayment: 15, LoanTermInMonths: 15},
}

func GetAllBanks() *[]Bank {
	return &bankList
}

func AddBank(add *Bank) {
	bankList = append(bankList, *add)
}

func UpdateBank(bankId int, updateBank *Bank) {
	for i, bank := range bankList {
		if bankId == bank.Id {
			bankList[i].Name = updateBank.Name
			bankList[i].InterestRate = updateBank.InterestRate
			bankList[i].MaximumLoan = updateBank.MaximumLoan
			bankList[i].MinimumDownPayment = updateBank.MinimumDownPayment
			bankList[i].LoanTermInMonths = updateBank.LoanTermInMonths
			break
		}
	}
}

func GetBankByID(id int) (*Bank, error) {
	for _, bank := range bankList {
		if bank.Id == id {
			return &bank, nil
		}
	}
	return nil, errors.New("Bank not found")
}

func DeleteBank(id int) string {
	var name string
	for index, bank := range bankList {
		if bank.Id == id {
			name = bank.Name
			bankList = append(bankList[:index], bankList[index+1:]...)
		}
	}
	return name
}
