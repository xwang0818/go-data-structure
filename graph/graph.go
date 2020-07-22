package graph

func containsCycleDirectedGraph(n int, edges [][]int) bool {
    dict := make(map[int][]int)
    for _, edge := range edges {
        dict[edge[1]] = append(dict[edge[1]], edge[0])
    }
    // 0 = unvisited, 1 = visited, 2 = processing
    visited := make([]int, n, n)
    for i := 0; i < n; i++ {
        if containsCycleDirectedGraphHelper(i, dict, visited) == false {
            return false
        }
    }
    return true
}
// DFS helper
func containsCycleDirectedGraphHelper(node int, dict map[int][]int, visited []int) bool {
    if visited[node] == 1 {
        return true
    }
    if visited[node] == 2 {
        return false
    }
    visited[node] = 2
    result := true
    for _, item := range dict[node] {
        result = result && containsCycleDirectedGraphHelper(item, dict, visited)
    }
    visited[node] = 1
    return result
}

func containsCycleUndirectedGraph(n int, connections [][]int) bool {
    dict := make(map[int][]int)
    for _, edge := range connections {
        dict[edge[1]] = append(dict[edge[1]], edge[0])
        dict[edge[0]] = append(dict[edge[0]], edge[1])
    }
    visited := make([]bool, n, n)
    for i := 0; i < n; i++ {
        if !visited[i] {
            if containsCycleUndirectedGraphHelper(i, dict, visited, -1) {
                return true
            }
        }
    }
    return false
}
// DFS helper
func containsCycleUndirectedGraphHelper(node int, dict map[int][]int, visited []bool, parent int) bool {
    visited[node] = true
    for _, neigh := range dict[node] {
        if !visited[neigh] {
            if containsCycleUndirectedGraphHelper(neigh, dict, visited, node) {
                return true
            }
        } else if neigh != parent {
            return true
        }
    }
    return false
}

func stronglyConnectedComponents(n int, edges[][]int) [][]int {
    adj := make(map[int][]int)
    reverse := make(map[int][]int)
    for _, edge := range edges {
        adj[edge[0]] = append(adj[edge[0]], edge[1])
        reverse[edge[1]] = append(reverse[edge[1]], edge[0])
    }
    visited := make([]bool, n)
    stack := []int{}
    for i := 0; i < n; i++ {
        stack = append(stack, stronglyConnectedComponentsHelper(i, adj, visited)...)
    }
    result := [][]int{}
    visited = make([]bool, n)
    for len(stack) != 0 {
        last := len(stack)-1
        node := stack[last]
        stack = stack[:last]
        comp := stronglyConnectedComponentsHelper(node, reverse, visited)
        if len(comp) != 0 {
            result = append(result, comp)
        }
    }
    return result
}
// DFS helper
func stronglyConnectedComponentsHelper(node int, adj map[int][]int, visited []bool) []int {
    if visited[node] {
        return []int{}
    }
    visited[node] = true
    result := []int{}
    for _, neigh := range adj[node] {
        result = append(result, stronglyConnectedComponentsHelper(neigh, adj, visited)...)
    }
    result = append(result, node)
    return result
}
