package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlightsCalculation(t *testing.T) {

	type given struct {
		flights []Flight
		src     string
		dst     string
	}
	type expected struct {
		flights int
	}
	var cases = []struct {
		name     string
		given    given
		expected expected
	}{
		{
			name: "Should return 1 flight",
			given: given{
				flights: []Flight{
					{"SFO", "EWR"},
				},
				src: "SFO",
				dst: "EWR",
			},
			expected: expected{
				flights: 1,
			},
		},
		{
			name: "Should return 2 flight",
			given: given{
				flights: []Flight{
					{"ATL", "EWR"},
					{"SFO", "ATL"},
				},
				src: "SFO",
				dst: "EWR",
			},
			expected: expected{
				flights: 2,
			},
		},
		{
			name: "Should return 4 flight",
			given: given{
				flights: []Flight{
					{"IND", "EWR"},
					{"SFO", "ATL"},
					{"GSO", "IND"},
					{"ATL", "GSO"},
				},
				src: "SFO",
				dst: "EWR",
			},
			expected: expected{
				flights: 4,
			},
		},
	}
	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			flights := FlightsCalculation(tc.given.flights, tc.given.src, tc.given.dst)
			assert.Equal(t, tc.expected.flights, flights)

		})

	}

}
