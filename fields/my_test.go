package fields

import "testing"

func TestCustomVsItem(t *testing.T) {
	s := Person{}
	ItemStructField(s)
	CustomStructField(s)
}
