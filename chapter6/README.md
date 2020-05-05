### Graph Breadth First Search  (BFS)

- Go through the graph
- Check if a value exists in the graph
- Find the shortest way (Doesn't work for the weighted graph)

Time comsuming: O(V + E) (V: number of vertice, E: number of edge)

Graph1 for problem if can find a person in facebook network
![alt text](https://github.com/fooxlj07/grokking-algorithms/blob/master/pictures/graph1.png?raw=true)

There is two ways to store the graph, as the picture showed.

1. hashtable + list
```
map[string][]string
you: [Bob, Claire, Alice]
Bob: [Peggy, Anuj]
Claire: [Thom, Jonny]
Alice: [Peggy]
Peggy: []
Anuj: []
Thom: [] 
Jonny: []
```

2. Matrix
```
        You Bob Claire Alice Peggy Anuj Thom Jonny
You      0   1    1      1    0     0     0    0
Bob      0   0    0      0    1     1     0    0
Claire   0   0    0      0    0     0     1    1
Alice    0   0    0      0    1     0     0    0
Peggy    0   0    0      0    0     0     0    0
Anuj     0   0    0      0    0     0     0    0
Thom     0   0    0      0    0     0     0    0
Jonny    0   0    0      0    0     0     0    0
```

3. Hashtable + linked list
```
you -> Bob -> Claire -> Alice
Bob -> Peggy -> Anuj
Claire -> Thom -> Jonny
Alice -> Peggy
Peggy -> 
Anuj ->
Thom -> 
Jonny ->
```
Optionals: Orthogonal List, Adjacency multiple table
