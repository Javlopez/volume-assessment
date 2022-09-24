package calculate

type FlightsParams struct {
	Flights []Flight `json:"flights"`
	Route   []string `json:"route"`
}

type FlightPaths struct {
	FlightPaths int           `json:"flight_paths"`
	Source      FlightsParams `json:"source"`
}

type Flight []string

type GraphFlights map[string][]string

func NewGraph() GraphFlights {
	return GraphFlights{}
}

func (g *GraphFlights) Has(src string) bool {
	if _, ok := (*g)[src]; ok {
		return true
	}
	return false
}

func (g *GraphFlights) AddFlight(src string) {
	(*g)[src] = []string{}
}

func (g *GraphFlights) Append(src string, value string) {
	(*g)[src] = append((*g)[src], value)
}
