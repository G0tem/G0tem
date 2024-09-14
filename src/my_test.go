package house_test

import (
	"testing"

	house "github.com/G0tem/G0tem/src"
)

func TestMyHouse(t *testing.T) {
	if house.MyHouse() != "YES!" {
		t.Fail()
	}
}
