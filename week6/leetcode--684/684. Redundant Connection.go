func findRedundantConnection(edges [][]int) []int {
    graph := make(map[int][]int)
    nodes_num := len(edges)
    degree := make([]int, nodes_num+1)
    visited := make([]bool, nodes_num+1)
    queue := make([]int ,0)
    for _, edge :=range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
        degree[edge[0]]++
        degree[edge[1]]++
    }
    
    // for node, edges := range graph {
    //     fmt.Print(node, ": ")
    //     for _, edge := range edges {
    //         fmt.Print(edge, " ")
    //     }
    //     fmt.Println()
    // }
    
    for i:=1; i<=nodes_num; i++ {
        if degree[i] == 1 {
            queue = append(queue, i)
        }
    }

    for len(queue) > 0 {
        head := queue[0]
        queue = queue[1:]
        visited[head] = true
        for _, edge := range graph[head] {
            degree[edge]--
            if degree[edge] == 1 {
                queue = append(queue, edge)
            }
        }
    }

    for i:=nodes_num-1; i>=0 ;i-- {
        if visited[edges[i][0]] == false && visited[edges[i][1]] == false {
            return edges[i]
        }
    }
    return nil
}