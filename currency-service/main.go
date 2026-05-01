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
			"status":           "success",
			"original_amount":  amount,
			"base_currency":    "KZT",
			"target_currency":  "USD",
			"exchange_rate":    exchangeRate,
			"converted_amount": fmt.Sprintf("%.2f", usdResult),
			"message":          "Converted by Currency Microservice",
		})
	})

	fmt.Println("Currency Service is running on port 8081...")
	r.Run(":8081")
}
