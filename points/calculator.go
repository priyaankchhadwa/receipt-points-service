package points

import (
	"math"
	"strconv"
	"strings"
	"time"

	"receipt-points-service/models"
)

func CalculatePoints(receipt models.Receipt) int {
	totalPoints := 0

	// 1. One point for every alphanumeric character in the retailer name
	totalPoints += countAlphanumericCharacters(receipt.Retailer)

	// 2. 50 points if total is a round dollar amount
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 1) == 0 {
		totalPoints += 50
	}

	// 3. 25 points if total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		totalPoints += 25
	}

	// 4. 5 points for every two items
	totalPoints += (len(receipt.Items) / 2) * 5

	// 5. Special points for item descriptions
	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			totalPoints += int(math.Ceil(itemPrice * 0.2))
		}
	}

	// 6. 6 points if purchase day is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		totalPoints += 6
	}

	// 7. 10 points if purchase time is between 2pm and 4pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
		totalPoints += 10
	}

	return totalPoints
}

func countAlphanumericCharacters(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') {
			count++
		}
	}
	return count
}
