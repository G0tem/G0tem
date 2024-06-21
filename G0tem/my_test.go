package g0tem_test

import (
	"testing"

	g0tem "github.com/G0tem/G0tem/G0tem"
)

func TestMyHouse(t *testing.T) {
	if g0tem.MyHouse() != "NO" {
		t.Fail()
	}
}
