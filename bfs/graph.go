package bfs

import "fmt"

//Breadth First Search  O(V + E), time V: number of vertice, E: number of edge
func checkPersonExist(graph map[string][]string, name string) bool {
	queue := graph["you"]
	checked := map[string]int{}
	for len(queue) > 0 {
		fmt.Println(queue)
		p := queue[0]
		queue = queue[1:]
		if _, ok := checked[p]; ok {
			continue
		} else {
			checked[p] = 1
		}

		if p != name {
			queue = append(queue, graph[p]...)
		} else {
			return true
		}
	}
	return false
}

func findShortestPath(graph map[string][]string, name string) []string {
	queue := []string{"you"}
	checked := map[string]int{}
	parent := map[string]string{}
	shortest := []string{name}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if _, ok := checked[p]; ok {
			continue
		} else {
			checked[p] = 1
		}

		if p != name {
			queue = append(queue, graph[p]...)
		} else {
			break
		}
		for _, a := range graph[p] {
			if _, ok := parent[a]; !ok {
				parent[a] = p
			}
		}
	}
	// means find the name
	p, ok := parent[name]
	for ok {
		shortest = append([]string{p}, shortest...)
		p, ok = parent[p]
	}
	if len(shortest) == 1 {
		return nil
	}
	return shortest
}

func graphInMapList() map[string][]string {
	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{"anuj"}
	graph["thom"] = []string{}
	graph["jonny"] = []string{"thom"}
	return graph
}

func Run() {
	fmt.Println(checkPersonExist(graphInMapList(), "thom"))
	fmt.Println(findShortestPath(graphInMapList(), "toto"))
}
