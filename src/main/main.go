package main

import (
	"github.com/lwfofgit/coding/src/algorithms/golang/search/bfs_dfs"
)

func main()  {

	g := bfs_dfs.CreateGraph(8)
	g.AddEdge(0,0)
	g.AddEdge(0,1)
	g.AddEdge(0,3)
	g.AddEdge(1,2)
	g.AddEdge(1,4)
	g.AddEdge(2,5)
	g.AddEdge(3,4)
	g.AddEdge(4,5)
	g.AddEdge(4,6)
	g.AddEdge(5,7)
	g.AddEdge(6,7)

	g.BFS(0,7)
}
