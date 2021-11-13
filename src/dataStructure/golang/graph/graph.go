package graph

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
	return &Graph{
		VertexNums:     v,
		AdjacencyTable: make([][]int, 0),
	}
}

func (g *Graph) AddEdge(s, v int) {
	g.AdjacencyTable[s] = append(g.AdjacencyTable[s], v)
	g.AdjacencyTable[v] = append(g.AdjacencyTable[v], s)
}
