package main

import "testing"

// Testing commands 
// go test -v
// go test -cover 
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out 

// AUTO TESTING
type Test struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}

var tests = []Test{
	{"Good Test", 10.0, 2.0, 5.0, false},
	{"Bad Test", 1000.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else if !tt.isErr {
			if err != nil {
				t.Error("Did not expect an Error but got one, life is hard", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}

}

// Here u must manually create a function for each test which is really inefficient
// This is the manual way where u create 1 test for each thing u want to test, and
// is really inefficient, or there are just better options now.
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
