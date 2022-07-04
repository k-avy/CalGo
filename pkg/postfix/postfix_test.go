package postfix

import "testing"


func TestSolvePostFixExpression(t *testing.T) {
	got, err := SolvePostfixExpression("4 6 +")
	want := 10
	if got != want || err != nil {
		t.Errorf("Error Found %v",err )
	}
}

func TestConvertPosttoPre(t *testing.T) {
	got, err := ConvertPosttoPre("4 6 +")
	want := "+ 4 6"
	if got != want || err != nil {
		t.Errorf("Error Found %v",err )
	}
}

func TestConvertPosttoIn(t *testing.T) {
	got, err := ConvertPosttoIn("4 6 +")
	want := "4 + 6"
	if got != want || err != nil {
		t.Errorf("Error Found %v",err )
	}
}
