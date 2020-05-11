### Dijkstra

Doens't work for Graph with negative weights (should use Bellman-Ford algorithm)
Doesn't work for cyclic graph，DAG

Algo 

1. Get the costest lowest node from the starting node
2. Calculate all the childs of lowest node (the `costs[lowestCostNode] + the weight child's node`)
3. If the costs[child node] have already exist, make `costs[child node] = min( costs[child node], costs[lowestCostNode] + the weight child's node)`
4. Repeat the first step

Solution for corresponding graph: 
![alt text](https://github.com/fooxlj07/grokking-algorithms/blob/master/pictures/dijkstra.png?raw=true)