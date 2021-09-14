package handler

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/entity"
)

func (h *Handler) showIndexPage(c *gin.Context) {
	banks := entity.GetAllBanks()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": banks,
		},
	)
}

func (h *Handler) getBank(c *gin.Context) {
	if bankId, err := strconv.Atoi(c.Param("bank_id")); err == nil {
		if bank, err := entity.GetBankByID(bankId); err == nil {
			c.HTML(
				http.StatusOK,
				"bank.html",
				gin.H{
					"title":   bank.Name,
					"payload": bank,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (h *Handler) createBank(c *gin.Context) {
	banks := entity.GetAllBanks()
	if c.Request.Method != http.MethodPost {
		c.HTML(
			http.StatusOK,
			"createBank.html",
			gin.H{
				"title": "Create bank",
			},
		)
		return
	} else {
		bankDetails := entity.Bank{}
		if c.Bind(&bankDetails) != nil {
			c.HTML(
				http.StatusInternalServerError,
				"index.html",
				gin.H{
					"title":   "Home page",
					"message": "Bank wasn't created. Server Error !!!",
					"payload": banks,
				},
			)
		} else {
			entity.AddBank(&bankDetails)
			c.HTML(
				http.StatusOK,
				"index.html",
				gin.H{
					"title":   "Home page",
					"message": "Bank success created !!!",
					"payload": banks,
				},
			)
		}
	}
}

func (h *Handler) updateBank(c *gin.Context) {
	if bankId, err := strconv.Atoi(c.Param("bank_id")); err == nil {
		if bank, err := entity.GetBankByID(bankId); err == nil {
			if c.Request.Method != http.MethodPost {
				c.HTML(
					http.StatusOK,
					"updateBank.html",
					gin.H{
						"title": "Update " + bank.Name,
						"bank":  bank,
					},
				)
			} else {
				bankDetails := entity.Bank{}
				if c.Bind(&bankDetails) == nil {
					entity.UpdateBank(bankId, &bankDetails)
					banks := entity.GetAllBanks()
					c.HTML(
						http.StatusOK,
						"index.html",
						gin.H{
							"title":   "Home page",
							"message": "Bank " + bankDetails.Name + " success updated!!!",
							"payload": banks,
						},
					)
				}
			}
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (h *Handler) deleteBank(c *gin.Context) {
	banks := entity.GetAllBanks()
	if bankId, err := strconv.Atoi(c.Param("bank_id")); err == nil {
		bankName := entity.DeleteBank(bankId)
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title":   "Home page",
				"message": bankName + " successful delete !!!",
				"payload": banks,
			},
		)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

type inputData struct {
	InitialLoan float64 `form:"InitialLoan"`
	DownPayment float64 `form:"DownPayment"`
	BankId      int     `form:"bank"`
}

func (h *Handler) MortgageCalculator(c *gin.Context) {
	banks := entity.GetAllBanks()
	if c.Request.Method != http.MethodPost {
		c.HTML(
			http.StatusOK,
			"MortgageCalculator.html",
			gin.H{
				"title":   "Mortgage Calculator",
				"payload": banks,
			},
		)
	} else {
		mortgageDetails := inputData{}
		if c.Bind(&mortgageDetails) != nil {
			c.HTML(
				http.StatusInternalServerError,
				"index.html",
				gin.H{
					"title":   "Mortgage Calculator",
					"message": "Entry data again, please. Server Error !!!",
					"payload": banks,
				},
			)
		} else {
			bank, err := entity.GetBankByID(mortgageDetails.BankId)
			if bank.MaximumLoan >= mortgageDetails.InitialLoan && bank.MinimumDownPayment <= mortgageDetails.DownPayment && err == nil {
				mortgage := (bank.InterestRate / 12) * math.Pow((1+bank.InterestRate/12), float64(bank.LoanTermInMonths))
				mortgage = mortgage / (math.Pow((1+bank.InterestRate/12), float64(bank.LoanTermInMonths)) - float64(1))
				c.HTML(
					http.StatusOK,
					"MortgageCalculator.html",
					gin.H{
						"title":   "Mortgage Calculator",
						"message": fmt.Sprintf("Your result: %f", mortgage),
						"payload": banks,
					},
				)
			} else {
				var message string
				if bank.MaximumLoan < mortgageDetails.InitialLoan {
					message += fmt.Sprintf("Error: Maximum Loan.The bank (%s) isn't capable of giving a requested loan.", bank.Name)
				}
				if bank.MinimumDownPayment > mortgageDetails.DownPayment {
					message += fmt.Sprintf("Error: down payment does not satisfy the minimum down payment boundary of the bank (%s).", bank.Name)
				}
				c.HTML(
					http.StatusOK,
					"MortgageCalculator.html",
					gin.H{
						"title":   "Mortgage Calculator",
						"message": message,
						"payload": banks,
					},
				)
			}
		}
	}
}
