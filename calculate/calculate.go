package calculate

func FlightsCalculation(flights []Flight, src, dst string) int {
	// Order the edges: convert flights into a graph
	graph := buildGraph(flights)
	visited := NewSet()
	visited.Add(src)

	q := NewQueue()
	q.Push(Node{Name: src, Distance: 0})
	for {
		current := q.Shift()
		node := current.Name
		distance := current.Distance

		if node == dst {
			return distance
		}

		for _, neighbor := range graph[node] {
			if !visited.Has(neighbor) {
				visited.Add(neighbor)
				nodeToAdd := Node{
					Name:     neighbor,
					Distance: distance + 1,
				}
				q.Push(nodeToAdd)
			}
		}

		if q.IsEmpty() {
			break
		}
	}
	return 0
}

func buildGraph(flights []Flight) GraphFlights {
	graph := NewGraph()
	for _, edge := range flights {
		src := edge[0]
		dst := edge[1]

		if !graph.Has(src) {
			graph.AddFlight(src)
		}

		if !graph.Has(dst) {
			graph.AddFlight(dst)
		}

		graph.Append(src, dst)
		graph.Append(dst, src)
	}
	return graph
}
