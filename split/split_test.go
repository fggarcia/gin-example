package split

import (
    "testing"
)

func TestSplitBy(t *testing.T) {
	result1, result2 := splitByDashUsingSplit("12-34", '-')
    if result1 != "12" || result2 != "34" {
        t.Error("Expected '12-34' to be split into '12' and '34'")
    }

    result3, result4 := splitByDashUsingBytes("12-34", '-')
    if result3 != "12" || result4 != "34" {
        t.Error("Expected '12-34' to be split into '12' and '34'")
    }
}
