package test

import (
  "testing"
  "e-mar404/receipt-processor/models"
)


func test_rule (t *testing.T, ruleFunction func(*models.Receipt) int) {
  receipt1 := models.Receipt {
    Retailer: "Target",
    PurchaseDate: "2022-01-01",
    PurchaseTime: "13:01",
    Total: "35.35",
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
    Retailer: "M&M Corner Market",
    PurchaseDate: "2202-03-20",
    PurchaseTime: "14:33",
    Total: "9.00",
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

  tests := []struct {
		expected  int
    receipt   models.Receipt
	}{
    {
      expected: 28,
      receipt: receipt1, 
    },
		{
      expected: 109,
      receipt: receipt2,
    },
	}

	for _, tt := range tests {
    t.Run("Count Points From Rules", func(t *testing.T) {
      if got := ruleFunction(&tt.receipt); got != tt.expected {
        t.Errorf("points counted:  %v, expected %v", got, tt.expected)
			}
		})
	}
}
