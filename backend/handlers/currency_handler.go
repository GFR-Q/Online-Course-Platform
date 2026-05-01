package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func GetConvertedPrice(c *gin.Context) {
	amount := c.Query("amount")
	if amount == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр amount обязателен"})
		return
	}

	client := resty.New()

	resp, err := client.R().
		SetQueryParam("amount", amount).
		Get("http://localhost:8081/convert")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Второй сервис (Currency) недоступен"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}
