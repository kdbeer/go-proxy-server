package main

import "testing"

type TC struct {
	A      float64
	B      float64
	Ops    string
	Expect float64
	ERROR  string
}

func TestSuccessCase(t *testing.T) {
	tt := []TC{
		{100, 10, "/calculator.sum", 110, ""},
		{100, 10, "/calculator.sub", 90, ""},
		{100, 10, "/calculator.div", 10, ""},
		{100, 10, "/calculator.mul", 1000, ""},
	}

	for _, tc := range tt {
		t.Run(
			tc.Ops, func(t *testing.T) {
				res, err := calculate(tc.A, tc.B, tc.Ops)
				if err != nil {
					t.Error("expect success but got error")
				}

				if res != tc.Expect {
					t.Errorf("expect %f ebot got %f", tc.Expect, res)
				}
			},
		)
	}
}

func TestCaseSouldError(t *testing.T) {
	tt := []TC{
		TC{100, 0, "/calculator.div", 0, "cannot divide by zero"},
	}

	for _, tc := range tt {
		t.Run(
			tc.Ops, func(t *testing.T) {
				_, err := calculate(tc.A, tc.B, tc.Ops)

				if err == nil {
					t.Error("expect error but got success")
					return
				}

				if err.Error() != tc.ERROR {
					t.Errorf("expect error %s ebot got %s", tc.ERROR, err)
				}
			},
		)
	}
}
