package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestApi(t *testing.T) {
	type given struct {
		url  string
		body string
		src  string
		dst  string
	}
	type expected struct {
		body       string
		statusCode int
	}
	var cases = []struct {
		name     string
		given    given
		expected expected
	}{
		{
			name: "Should raise error if path is wrong",
			given: given{
				url: "/message",
			},
			expected: expected{
				statusCode: http.StatusNotFound,
			},
		},
		{
			name: "Should raise error if missing body data",
			given: given{
				url: "/calculate",
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Should raise error if the json is malformed",
			given: given{
				url:  "/calculate",
				body: "[hey",
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Should raise error if the route is malformed",
			given: given{
				url:  "/calculate",
				body: `{"flights": [["SFO", "EWR"]], "route": ["SFO"]}`,
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "Should return 1 flight_paths",
			given: given{
				url:  "/calculate",
				body: `{"flights": [["SFO", "EWR"]], "route": ["SFO", "EWR"]}`,
			},
			expected: expected{
				statusCode: http.StatusOK,
				body:       `{"source": {"flights": [["SFO", "EWR"]], "route": ["SFO", "EWR"]}, "flight_paths":1}`,
			},
		},
		{
			name: "Should return 2 flight_paths",
			given: given{
				url:  "/calculate",
				body: `{"flights": [["ATL", "EWR"], ["SFO", "ATL"]] , "route": ["SFO", "EWR"]}`,
			},
			expected: expected{
				statusCode: http.StatusOK,
				body:       `{"source": {"flights": [["ATL", "EWR"], ["SFO", "ATL"]], "route": ["SFO", "EWR"]}, "flight_paths":2}`,
			},
		},
		{
			name: "Should return 4 flight_paths",
			given: given{
				url:  "/calculate",
				body: `{"flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]], "route": ["SFO", "EWR"]}`,
			},
			expected: expected{
				statusCode: http.StatusOK,
				body:       `{"source": {"flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]], "route": ["SFO", "EWR"]}, "flight_paths":4}`,
			},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			apitest.New().
				Handler(NewApi().Router).
				Post(tc.given.url).
				Body(tc.given.body).
				Expect(t).
				Status(tc.expected.statusCode).
				Body(tc.expected.body).
				End()

		})
	}
}
