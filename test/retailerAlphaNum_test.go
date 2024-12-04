package test

import (
  "testing"
	"e-mar404/receipt-processor/rules"
)

func Test_point_count_from_rules(t *testing.T) {
  test_rule(t, rules.CountPoints)
}
