package rules

import (
	"e-mar404/receipt-processor/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func retailerAlphaNumPoints(receipt *models.Receipt) int {
  points := 0

  for _, letter := range receipt.Retailer {
    if unicode.IsLetter(letter) || unicode.IsDigit(letter) {
      points += 1
    }
  }
  
  fmt.Printf("retail alpha: %v\n", points)

  return points
}

func roundDollarPoints(receipt *models.Receipt) int {
	points := 0

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return points
	}

	if total == float64(int(total)) {
		points = 50
	}

  fmt.Printf("round dollars: %v\n", points)
	return points
}

func quarterMultiplePoints(receipt *models.Receipt) int {
	points := 0

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return points
	}
  
	if (int(total * 100) % 25) == 0 {
		points = 25
	}

  fmt.Printf("quarter multiple: %v\n", points)
	return points
}

func multipleItemsPoints(receipt *models.Receipt) int {
  points := (len(receipt.Items) / 2) * 5

  fmt.Printf("multiple items: %v\n", points)

  return points
}

func descriptionLengthPoints(receipt *models.Receipt) int {
	points := 0

	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.Description)

		if len(trimmedDescription)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				continue
			}

			points += int(math.Ceil(price * 0.2))
		}
	}

  fmt.Printf("description length: %v\n", points)
	return points
}

func oddDayPoints(receipt *models.Receipt) int {
	points := 0

	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return points
	}

	if date.Day()%2 != 0 {
		points = 6
	}

  fmt.Printf("odd day: %v\n", points)
	return points
}

// messed up
func timeRangePoints(receipt *models.Receipt) int {
	points := 0

	timeParsed, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
    fmt.Printf("time range: %v\n", points)
		return points
	}

	startTime := time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)
	endTime := time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)

	if timeParsed.After(startTime) && timeParsed.Before(endTime) {
		points = 10
	}

	return points
}

func CountPoints(receipt *models.Receipt) int {
	points := 0

	points += retailerAlphaNumPoints(receipt)
	points += roundDollarPoints(receipt)
	points += quarterMultiplePoints(receipt)
	points += multipleItemsPoints(receipt)
	points += descriptionLengthPoints(receipt)
	points += oddDayPoints(receipt)
	points += timeRangePoints(receipt)

	return points
}
