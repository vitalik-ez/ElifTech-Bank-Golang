package handler

import (
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
	if bankId, err := strconv.Atoi(c.Param("bank_id")); err == nil {
		bankName := entity.DeleteBank(bankId)
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title":   "Home page",
				"message": bankName + " successful delete !!!",
			},
		)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
