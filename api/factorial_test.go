package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	api, err := Construct()
	if err != nil {
		t.Errorf("Error is: %v . Expected: %v", err, nil)
	}
	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: "4", output: "24" + "\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("GET", "/factorial/calculate?number="+test.input, strings.NewReader(""))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(api.CalculateHandler)

		handler.ServeHTTP(rr, req)

		if rr.Result().StatusCode != test.status {
			t.Errorf("Response status is: %v . Expected: %v", rr.Result().StatusCode, test.status)
		}

		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != string(test.output) {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}

	}
	Destruct(api.Domain.Pkg.Repository.Database)
}
