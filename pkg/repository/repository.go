package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllBanks() ([]entity.Bank, error) {
	var banks []entity.Bank
	rows, err := r.db.Query("SELECT * FROM banks")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		b := entity.Bank{}
		err := rows.Scan(&b.Id, &b.Name, &b.InterestRate, &b.MaximumLoan, &b.MinimumDownPayment, &b.LoanTermInMonths)
		if err != nil {
			fmt.Println(err)
			continue
		}
		banks = append(banks, b)
	}
	return banks, err
}

func (r *Repository) CreateBank(bank entity.Bank) error {
	_, err := r.db.Exec("INSERT INTO banks(name, interestrate, maximumloan, minimumdownpayment, loanterminmonths) VALUES ($1, $2, $3, $4, $5)", bank.Name, bank.InterestRate, bank.MaximumLoan, bank.MinimumDownPayment, bank.LoanTermInMonths)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (r *Repository) UpdateBank(bankId int, bank *entity.Bank) error {
	_, err := r.db.Exec("update banks set name = $1, interestrate = $2, maximumloan = $3, minimumdownpayment = $4, loanterminmonths = $5 where id = $6", bank.Name, bank.InterestRate, bank.MaximumLoan, bank.MinimumDownPayment, bank.LoanTermInMonths, bankId)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetBankByID(bankId int) (*entity.Bank, error) {
	row := r.db.QueryRow("SELECT * FROM banks WHERE id = $1", bankId)
	b := entity.Bank{}
	err := row.Scan(&b.Id, &b.Name, &b.InterestRate, &b.MaximumLoan, &b.MinimumDownPayment, &b.LoanTermInMonths)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &b, err
}

func (r *Repository) DeleteBank(bankId int) (string, error) {
	var name string
	row := r.db.QueryRow("DELETE FROM banks WHERE id = $1 RETURNING name", bankId)
	err := row.Scan(&name)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return name, nil
}
