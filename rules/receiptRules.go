package rules

import (
	"math"
	"time"
  "strconv"
	"unicode"
  "strings"
  "e-mar404/receipt-processor/models"
)

func RetailerAlphaNumPoints(receipt *models.Receipt) int {
  points := 0

  for _, letter := range receipt.Retailer {
    if unicode.IsLetter(letter) || unicode.IsDigit(letter) {
      points += 1
    }
  }
  
  return points
}

func RoundDollarPoints(receipt *models.Receipt) int {
	points := 0

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return points
	}

	if total == float64(int(total)) {
		points = 50
	}

	return points
}

func QuarterMultiplePoints(receipt *models.Receipt) int {
	points := 0

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return points
	}
  
	if (int(total * 100) % 25) == 0 {
		points = 25
	}

	return points
}

func MultipleItemsPoints(receipt *models.Receipt) int {
  return (len(receipt.Items) / 2) * 5
}

func DescriptionLengthPoints(receipt *models.Receipt) int {
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

	return points
}

func OddDayPoints(receipt *models.Receipt) int {
	points := 0

	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return points
	}

	if date.Day()%2 != 0 {
		points = 6
	}

	return points
}

func TimeRangePoints(receipt *models.Receipt) int {
	points := 0

	timeParsed, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
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

	points += RetailerAlphaNumPoints(receipt)
	points += RoundDollarPoints(receipt)
	points += QuarterMultiplePoints(receipt)
	points += MultipleItemsPoints(receipt)
	points += DescriptionLengthPoints(receipt)
	points += OddDayPoints(receipt)
	points += TimeRangePoints(receipt)

	return points
}
