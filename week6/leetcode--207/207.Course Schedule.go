func findCycle(s int, visited []int, graph map[int][]int) bool {
    visited[s] = -1
    for _, edge := range graph[s] {
        if visited[edge] == -1 {
            return true
        } else if visited[edge] == 0 {
            if findCycle(edge, visited, graph) {
                return true
            }
        }
    }
    visited[s] = 1
    return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make(map[int][]int)
    visited := make([]int, numCourses)

    for _, pair := range prerequisites {
        graph[pair[1]] = append(graph[pair[1]], pair[0])
    }
    for i:=0; i<numCourses; i++ {
        if visited[i] == 0 {
            if findCycle(i, visited, graph) {
                return false
            }
        }
    }
    return true
}