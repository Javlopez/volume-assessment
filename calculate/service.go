package calculate

import "fmt"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) FlightPaths(params FlightsParams) (int, error) {

	if len(params.Route) != 2 {
		return 0, fmt.Errorf("unable to calculate the route")
	}

	src := params.Route[0]
	dst := params.Route[1]

	return FlightsCalculation(params.Flights, src, dst), nil
}
