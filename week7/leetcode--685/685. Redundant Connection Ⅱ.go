func findRedundantDirectedConnection(edges [][]int) []int {
    graph := make(map[int]int)
    graph2 := make(map[int][]int)
    nodes_count := len(edges)
    in_degree := make([]int, nodes_count+1)
    double_in_degree := -1
    candidate := -1
    for _, edge := range edges {
        if in_degree[edge[1]] > 0 {
            double_in_degree = edge[1]
            candidate = edge[0]
        } else {
            graph2[edge[0]] = append(graph2[edge[0]], edge[1])
            graph[edge[1]] = edge[0]
            in_degree[edge[1]]++
        }
    }

    if double_in_degree == -1 {
        parents := make([]int ,len(edges)+1)
        for _, edge := range edges {
            parent0 := findParent(edge[0], parents)
            parent1 := findParent(edge[1], parents)

            if parent0 == parent1 {
                return edge
            }
            parents[parent1] = parent0
        }
    } else {
        // for key, value := range graph2 {
        //     fmt.Println(key, value)
        // }
        queue := make([]int, 0)
        visited := make([]int, 0)

        for i:=1; i<=nodes_count; i++ {
            if in_degree[i] == 0 {
                queue = append(queue, i)
                visited = append(visited, i)
                // fmt.Println(i)
            }
        }

        if len(queue) > 1 {
            return []int{graph[double_in_degree], double_in_degree}
        }

        for len(queue)!=0 {
            head := queue[0]
            queue = queue[1:]
            for _, next := range graph2[head] {
                in_degree[next]--
                if in_degree[next] ==0 {
                    queue = append(queue, next)
                    visited = append(visited, next)
                }
            }
        }
        if len(visited) == nodes_count {
            return []int{candidate, double_in_degree}
        }
        return []int{graph[double_in_degree], double_in_degree}
    }
    return []int{}
}

func findParent(node int, parents []int) int {
    if parents[node] == 0 {
        return node
    }
    return findParent(parents[node], parents)
}