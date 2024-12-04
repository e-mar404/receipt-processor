package test

import (
  "testing"
	"e-mar404/receipt-processor/rules"
  "e-mar404/receipt-processor/models"
)

func Test_time_range(t *testing.T) {
  ruleName := "Purchase time is between range (2-4pm)"

  receipt1 := models.Receipt {
    PurchaseTime: "13:01",
  }

  receipt2 := models.Receipt {
    PurchaseTime: "14:33",
  }

  test_rule(ruleName, t, rules.TimeRangePoints, receipt1, 0)
  test_rule(ruleName, t, rules.TimeRangePoints, receipt2, 10)
}
