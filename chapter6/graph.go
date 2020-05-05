package chapter6

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

func graphInMapList() map[string][]string {
	graph := map[string][]string{}
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	return graph
}

func Run() {
	fmt.Println(checkPersonExist(graphInMapList(), "toto"))
}
