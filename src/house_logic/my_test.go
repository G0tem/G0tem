package house_test

import (
	"testing"

	house "github.com/G0tem/G0tem/src/house_logic"
)

func TestMyHouse(t *testing.T) {
	if house.MyHouse() != "YES!" {
		t.Fail()
	}
}
