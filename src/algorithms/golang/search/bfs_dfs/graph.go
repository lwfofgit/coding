package bfs_dfs

/*
图表示方法：
 1. 邻接矩阵
 2. 邻接表
*/

type Graph struct {
	VertexNums     int 		// 顶点个数
	AdjacencyTable [][]int // 邻接表
}

func CreateGraph(v int) *Graph {
	a := make([][]int, 8)
	for i := range a {
		a[i] = make([]int, 0)
	}
	return &Graph{
		VertexNums:     v,
		AdjacencyTable: a,
	}
}

func (g *Graph) AddEdge(s, v int) {
	g.AdjacencyTable[s] = append(g.AdjacencyTable[s], v)
	g.AdjacencyTable[v] = append(g.AdjacencyTable[v], s)
}
