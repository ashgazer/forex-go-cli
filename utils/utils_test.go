package utils

import "testing"

func TestCalcChange(t *testing.T) {
	want := 20.00
	if got := CalcChange(10.0, 2.0); got != want {

		t.Errorf("CalcChange(10.0, 20.0 = %f want %f", got, want)
	}
}

func TestRound(t *testing.T) {
	want := 3.325

	if got := round(3.32518634, 3); got != want {
		t.Errorf("round(3.32578634, 3) = %f want %f", got, want)
	}

}

func TestGetXURL(t *testing.T) {
	want := "https://api.ratesapi.io/api/latest?base=AOC&rtype=fpy&symbols=ABC"

	if got := getXURL("AOC", "ABC"); got != want {
		t.Errorf("getXURL(\"AOC\", \"ABC\") = %s want %s", got, want)
	}
}
