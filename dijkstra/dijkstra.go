package dijkstra

import "fmt"

func dijkstra(start string, end string, graph map[string]map[string]int) ([]string, int) {
	// To store the lowest costs to each node
	costs := map[string]int{}
	// To store the parent of the node
	parents := map[string]string{}
	// To use to check if from the to the node with lowest cost path node already calculcated. WHY?
	shortestPath := []string{end}
	for k, v := range graph[start] {
		costs[k] = v
		parents[k] = start
	}
	calculated := map[string]bool{start: true}
	lowestCostNode := findTheLowestCost(costs, calculated)
	for lowestCostNode != "" {
		calculated[lowestCostNode] = true
		for k, v := range graph[lowestCostNode] {
			if _, ok := costs[k]; ok && costs[k] < costs[lowestCostNode]+v {
				continue
			}
			costs[k] = costs[lowestCostNode] + v
			parents[k] = lowestCostNode
		}
		lowestCostNode = findTheLowestCost(costs, calculated)
	}
	parent, ok := parents[end]
	for ok {
		shortestPath = append([]string{parent}, shortestPath...)
		parent, ok = parents[parent]
	}
	return shortestPath, costs[end]
}

func findTheLowestCost(costs map[string]int, calculated map[string]bool) string {
	// Find the lowest cost in the costs
	lowest, node := 100000000, ""
	for k, v := range costs {
		if v <= lowest && !calculated[k] {
			lowest = v
			node = k
		}
	}
	return node
}

func buildGraph() map[string]map[string]int {
	graph := map[string]map[string]int{
		"A": map[string]int{"B": 5, "C": 0},
		"B": map[string]int{"D": 15, "E": 20},
		"C": map[string]int{"D": 10, "E": 35},
		"D": map[string]int{"F": 20},
		"E": map[string]int{"F": 10},
	}

	return graph
}

func Run() {
	fmt.Println(dijkstra("A", "F", buildGraph()))
}
