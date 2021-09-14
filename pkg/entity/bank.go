package entity

type Bank struct {
	Id                 int     `json:"id"`
	Name               string  `form:"Name"`
	InterestRate       float64 `form:"InterestRate" `
	MaximumLoan        float64 `form:"MaximumLoan"`
	MinimumDownPayment float64 `form:"MinimumDownPayment"`
	LoanTermInMonths   uint    `form:"LoanTermInMonths"`
}
