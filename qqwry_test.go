package qqwry

import (
	"testing"
)

func TestQueryIP(t *testing.T) {
	qqwry, err := NewQQWry("./QQWry.dat")
	if err != nil {
		t.Fatal(err)
	}

	country, area := qqwry.QueryIP("8.8.8.8")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("country:%s area:%s", country, area)
}
