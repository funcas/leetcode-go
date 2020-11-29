package problems

/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	g := newGraph(numCourses)
	for i := 0; i < len(prerequisites); i++ {
		g.addEdge(prerequisites[i][1], prerequisites[i][0])
	}
	ret := g.topoSortByKahn()
	if len(ret) == numCourses {
		return ret
	}
	return nil
}

type Graph struct {
	V   int
	Adj [][]int
}

func newGraph(v int) *Graph {
	t := make([][]int, v)
	return &Graph{
		V:   v,
		Adj: t,
	}
}

func (graph *Graph) addEdge(s, t int) {
	graph.Adj[s] = append(graph.Adj[s], t)
}

func (graph *Graph) topoSortByKahn() []int {
	var ans []int
	inDegree := make([]int, graph.V)
	for i := 0; i < graph.V; i++ {
		for j := 0; j < len(graph.Adj[i]); j++ {
			w := graph.Adj[i][j]
			inDegree[w]++
		}
	}
	var queue []int
	for i := 0; i < graph.V; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		ans = append(ans, i)
		for j := 0; j < len(graph.Adj[i]); j++ {
			k := graph.Adj[i][j]
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	return ans
}

// @lc code=end
