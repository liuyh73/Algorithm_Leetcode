func eventualSafeNodes(graph [][]int) []int {
	nodes := make([]int, 0)
	outDegree := make(map[int]int)
	prevNodes := make(map[int][]int)
	queue := make([]int, 0)
	for u, vs := range graph {
		outDegree[u] = len(vs)
		if outDegree[u] == 0 {
			queue = append(queue, u)
			nodes = append(nodes, u)
		}

		for _, v := range vs {
			prevNodes[v] = append(prevNodes[v], u)
		}
	}
	// fmt.Println(outDegree)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		for _, anotherNode := range prevNodes[node] {
			outDegree[anotherNode]--
			if outDegree[anotherNode] == 0 {
				nodes = append(nodes, anotherNode)
				queue = append(queue, anotherNode)
			}
		}
	}

	sort.Ints(nodes)

	return nodes
}