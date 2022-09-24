package docs

import "github.com/Javlopez/volume-assessment/calculate"

// swagger:route POST /calculate flights paramsFlights
// Calculate the flight paths
// responses:
//
//	200: apiCalculateSuccessResponse
//
// swagger:parameters paramsFlights
type apiCalculateParamsWrapper struct {
	// Execute a calculation of flights
	// in:body
	Body calculate.FlightsParams
}

// Successful calculation flights response.
// swagger:response apiCalculateSuccessResponse
type apiCalculateResponseWrapper struct {
	// in:body
	Body calculate.FlightPaths
}
