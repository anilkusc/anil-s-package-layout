package domain

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	d, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}

	tests := []struct {
		input  int
		output int
		err    error
	}{
		{input: 4, output: 24, err: nil},
	}
	for _, test := range tests {
		output, err := d.Calculate(test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if test.output != output {
			t.Errorf("Result is: %v . Expected: %v", output, test.output)
		}

	}

	Destruct(d.Pkg.Repository.Database)
}
