package bfs_dfs

import (
	"testing"
)

func TestGraph_BFS(t *testing.T) {
	//var buf bytes.Buffer
	//log.SetOutput(&buf)
	//defer func() {
	//	log.SetOutput(os.Stderr)
	//}()

	g := CreateGraph(8)
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
	//t.Log(buf.String())
}
