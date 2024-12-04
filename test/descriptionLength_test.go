package test

import (
  "testing"
	"e-mar404/receipt-processor/rules"
  "e-mar404/receipt-processor/models"
)

func Test_decription_length(t *testing.T) {
  ruleName := "Item description length"

  receipt1 := models.Receipt {
    Items: []models.Item {
      {
        Description: "Mountain Dew 12PK",
        Price: "6.49",
      },
      {
        Description: "Emils Cheese Pizza",
        Price: "12.25",
      },
      {
        Description: "Knorr Creamy Chicken",
        Price: "3.26",
      },
      {
        Description: "Doritos Nacho Cheese",
        Price: "3.35",
      },
      {
        Description: "Klarbrunn 12-PK 12 FL OZ",
        Price: "12.00",
      },
    },
  }

  receipt2 := models.Receipt {
    Items: []models.Item {
      {
        Description: "Gatorade",
        Price: "2.25",
      },
      {
        Description: "Gatorade",
        Price: "2.25",
      },
      {
        Description: "Gatorade",
        Price: "2.25",
      },
      {
        Description: "Gatorade",
        Price: "2.25",
      },
    },
  }

  test_rule(ruleName, t, rules.DescriptionLengthPoints, receipt1, 6)
  test_rule(ruleName, t, rules.DescriptionLengthPoints, receipt2, 0)
}
