package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/convert", func(c *gin.Context) {
		amountStr := c.Query("amount")
		if amountStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр amount обязателен"})
			return
		}

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат суммы"})
			return
		}

		exchangeRate := 500.0
		usdResult := amount / exchangeRate

		c.JSON(http.StatusOK, gin.H{
			"статус":            "success",
			"оригинал цена":     amount,
			"валюта":            "KZT",
			"перевод валюта":    "USD",
			"цена валюты":       exchangeRate,
			"сума при переводе": fmt.Sprintf("%.2f", usdResult),
			"оповещяние":        "Converted by Currency Microservice",
		})
	})

	fmt.Println("Currency Service is running on port 8081...")
	r.Run(":8081")
}
