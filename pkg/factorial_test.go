package pkg

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	p, f, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}

	tests := []struct {
		input  Factorial
		output int
		err    error
	}{
		{input: f, output: f.Result, err: nil},
	}
	for _, test := range tests {
		output, err := test.input.Calculate(p.Repository)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if test.output != output {
			t.Errorf("Result is: %v . Expected: %v", output, test.output)
		}

	}

	Destruct(p.Repository.Database)
}
