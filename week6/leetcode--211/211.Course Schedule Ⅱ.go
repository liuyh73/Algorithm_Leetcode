func canFinish(numCourses int, prerequisites [][]int) []int {
    graph := make(map[int][]int)
    in_degree := make([]int, numCourses)
    queue := make([]int, 0)
    schudule := make([]int, 0)
    for _, pair := range prerequisites {
        graph[pair[1]] = append(graph[pair[1]], pair[0])
        in_degree[pair[0]]++
    }

    for i:=0;i<numCourses;i++ {
        if in_degree[i] == 0 {
            queue = append(queue, i)
            schudule = append(schudule, i)
        }
    }

    for len(queue)!=0 {
        head := queue[0]
        queue = queue[1:]
        for _, edge := range graph[head] {
            in_degree[edge]--
            if in_degree[edge] == 0 {
                queue = append(queue, edge)
                schudule = append(schudule, edge)
            }
        }
    }
    return schudule
}