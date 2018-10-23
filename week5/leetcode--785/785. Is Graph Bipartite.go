func isBipartite(graph [][]int) bool {
    Graph := make(map[int][]int)
    var Color [150]int
    for index, _ := range Color {
        Color[index] = int(-1)
    }
    for node, edges := range graph {
        Graph[node] = edges
    }

    for index, _ :=range Color {
        if Color[index] == -1 {
            queue := make([]int, 0)
            queue = append(queue, index)
            head := 0
            Color[index] = 0
            for head < len(queue) {
                for _, edge := range Graph[queue[head]] {
                    if Color[edge] == -1 {
                        Color[edge] = 1 - Color[queue[head]]
                        queue = append(queue, edge)
                    } else if Color[edge] == Color[queue[head]] {
                        return false
                    }
                    // fmt.Println("node: ", edge, "color: ", Color[edge])
                }
                head++
                // fmt.Println(head)
            }
        }
    }
    return true
}