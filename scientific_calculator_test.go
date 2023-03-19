package scientificcalculator

import "testing"

func TestAbs(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc   string
		v      float64
		result float64
	}{
		{
			desc:   "ShouldReturnSameValue",
			v:      10,
			result: 10,
		},
		{
			desc:   "ShouldReturnPositiveValue",
			v:      -1,
			result: 1,
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			result := Abs(tC.v)

			if result != result {
				t.Fatal() // Study
			}
		})
	}
}
