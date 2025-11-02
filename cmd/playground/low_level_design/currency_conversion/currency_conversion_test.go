package currency_conversion

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		From     string
		To       string
		Amount   float64
		Rates    []Rate
		Expected float64
		Err      error
	}{
		{
			"BDT",
			"INR",
			1000,
			[]Rate{
				{
					"BDT",
					"USD",
					0.0082,
				},
				{
					"USD",
					"SGD",
					1.30,
				},
			},
			10.64,
			NoConversionFoundErr,
		},
	}

	for _, tt := range tests {
		got, err := Convert(tt.From, tt.To, tt.Amount, tt.Rates)
		if (got != tt.Expected) && (err != tt.Err) {
			t.Errorf("got: %v, expected %v\n", got, tt.Expected)
		}
	}
}
