package test

import (
  "testing"
	"e-mar404/receipt-processor/rules"
  "e-mar404/receipt-processor/models"
)

func Test_round_dollar(t *testing.T) {
  ruleName := "Total w Round Dollar"

  receipt1 := models.Receipt {
    Total: "35.35",
  }

  receipt2 := models.Receipt {
    Total: "9.00",
  }

  test_rule(ruleName, t, rules.RoundDollarPoints, receipt1, 0)
  test_rule(ruleName, t, rules.RoundDollarPoints, receipt2, 50)
}
