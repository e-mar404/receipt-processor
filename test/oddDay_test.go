package test

import (
  "testing"
	"e-mar404/receipt-processor/rules"
  "e-mar404/receipt-processor/models"
)

func Test_odd_day(t *testing.T) {
  ruleName := "Purchase date is odd"

  receipt1 := models.Receipt {
    PurchaseDate: "2022-01-01",
  }

  receipt2 := models.Receipt {
    PurchaseDate: "2022-03-20",
  }

  test_rule(ruleName, t, rules.OddDayPoints, receipt1, 6)
  test_rule(ruleName, t, rules.OddDayPoints, receipt2, 0)
}
