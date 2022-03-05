package domain

import (
	"errors"
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
		{input: -1, output: 0, err: errors.New("number cannot be smaller than 1")},
	}
	for _, test := range tests {
		output, err := d.Calculate(test.input)
		if err != test.err {
			if err.Error() != test.err.Error() {
				t.Errorf("Error is: %v . Expected: %v", err, test.err)
			}
		}
		if test.output != output {
			t.Errorf("Result is: %v . Expected: %v", output, test.output)
		}

	}

	Destruct(d.Pkg.Repository.Database)
}
