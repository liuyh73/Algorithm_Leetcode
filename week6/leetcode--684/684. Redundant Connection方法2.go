func findRedundantConnection(edges [][]int) []int {
    parents := make([]int ,len(edges)+1)

    for _, edge := range edges {
        parent0 := findParent(edge[0], parents)
        parent1 := findParent(edge[1], parents)

        if parent0 == parent1 {
            return edge
        }

        parents[parent0] = parent1
    }
    return []int{}
}

func findParent(node int, parents []int) int {
    if parents[node] == 0 {
        return node
    }
    return findParent(parents[node], parents)
}