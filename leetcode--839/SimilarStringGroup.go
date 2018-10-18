var graph = make(map[string][]string)
var visited = make(map[string]bool)
func numSimilarGroups(A []string) int {
	B := make([]string, 0)
	for _, str1 := range A {
		flag := false
		for _, str2 := range B {
			if str2 == str1 {
				flag = true
				break
			}
		}
		if !flag {
			B = append(B, str1)
		}
	}
	for i, str1 := range B {
		for j, str2 := range B {
            if i == j || str1 == str2 {
                continue
            }
            if length := len(str1); length == len(str2) {
				diffIndex := make([]int, 0)
                for k := 0; k < length; k++ {
					if str1[k] != str2[k] {
						diffIndex = append(diffIndex, k)
					}
				}
				if len(diffIndex) != 2 || !(str1[diffIndex[0]] == str2[diffIndex[1]] &&
					str1[diffIndex[1]] == str2[diffIndex[0]]) {
					continue
				}
				graph[str1] = append(graph[str1], str2)
			}
		}
	}

    num := 0
	for _, str := range B {
		if !visited[str] {
			num++
			visited[str] = true
			dfs(str)
		}
	}
	return num
}

func dfs(str string) {
	for _, strNext := range graph[str] {
		if !visited[strNext] {
			visited[strNext] = true
            dfs(strNext)
		}
	}
}