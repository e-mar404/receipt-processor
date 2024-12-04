package test

import (
  "testing"
	"e-mar404/receipt-processor/rules"
  "e-mar404/receipt-processor/models"
)

func Test_quarter_multiple(t *testing.T) {
  ruleName := "Total w quarter multiple"

  receipt1 := models.Receipt {
    Total: "35.35",
  }

  receipt2 := models.Receipt {
    Total: "9.00",
  }

  test_rule(ruleName, t, rules.QuarterMultiplePoints, receipt1, 0)
  test_rule(ruleName, t, rules.QuarterMultiplePoints, receipt2, 25)
}
