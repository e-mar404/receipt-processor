package test

import (
  "testing"
	"e-mar404/receipt-processor/rules"
  "e-mar404/receipt-processor/models"
)

func Test_retail_alpha_num(t *testing.T) {
  ruleName := "Retailer Alpha Num"
  receipt1 := models.Receipt {
    Retailer: "Target",
  }

  receipt2 := models.Receipt {
    Retailer: "M&M Corner Market",
  }

  test_rule(ruleName, t, rules.RetailerAlphaNumPoints, receipt1, 6)
  test_rule(ruleName, t, rules.RetailerAlphaNumPoints, receipt2, 14)
}
