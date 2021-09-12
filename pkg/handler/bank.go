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

}
