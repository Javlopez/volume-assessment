package calculate

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var params FlightsParams

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		errorResponse(Error{Error: "the request is wrong"}, w)
		return
	}

	svc := NewService()
	totalPaths, err := svc.FlightPaths(params)
	if err != nil {
		errorResponse(Error{Error: err.Error()}, w)
		return
	}

	response := FlightPaths{
		Source: FlightsParams{
			Flights: params.Flights,
			Route:   params.Route,
		},
		FlightPaths: totalPaths,
	}

	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(response)
	_, err = w.Write(bytes)
	if err != nil {
		panic(err)
	}
}

func errorResponse(e Error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	bytes, _ := json.Marshal(e)
	_, err := w.Write(bytes)
	if err != nil {
		panic(err)
	}
}
