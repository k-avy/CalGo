package prefix

import "testing"

func TestSolvePostFixExpression (t *testing.T){
	got,err:=SolvePrefixExpression("+ 4 6")
	want :=10
	if got!=want ||err!=nil {
			t.Errorf("error found")
	}
}

func TestConvertPosttoPre (t *testing.T){
	got,err:=ConvertPretoPost("+ 4 6")
	want:="4 6 +"
	if got!=want ||err!=nil {
		t.Errorf("error found")
}
}

func TestConvertPosttoIn (t *testing.T){
	got,err:=ConvertPretoIn("+ 4 6")
	want:="4 + 6"
	if got!=want ||err!=nil {
		t.Errorf("error found")
}
}