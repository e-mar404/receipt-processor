package test

import (
  "testing"
  "e-mar404/receipt-processor/models"
)

func test_rule (ruleName string, t *testing.T, ruleFunction func(*models.Receipt) int, receipt models.Receipt, expected int) {
  t.Run(ruleName, func(t *testing.T) {
    if got := ruleFunction(&receipt); got != expected {
      t.Errorf("points counted:  %v, expected %v", got, expected)
    }
  })
}
