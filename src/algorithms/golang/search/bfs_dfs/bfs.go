package bfs_dfs

import (
	"fmt"
)

// 图的广度优先
// prev: 记录搜索路径，prev[w]存储的是，顶点 w 是从哪个前驱顶点遍历过来的。
// 比如，我们通过顶点 2 的邻接表访问到顶点 3，那 prev[3]就等于 2
// visited:记录已经访问过的顶点
// queue: 是一个队列，用来存储已经被访问、但相连的顶点还没有被访问的顶点
/*

public void bfs(int s, int t) {
  if (s == t) return;
  boolean[] visited = new boolean[v];
  visited[s]=true;
  Queue<Integer> queue = new LinkedList<>();
  queue.add(s);
  int[] prev = new int[v];
  for (int i = 0; i < v; ++i) {
    prev[i] = -1;
  }
  while (queue.size() != 0) {
    int w = queue.poll();
   for (int i = 0; i < adj[w].size(); ++i) {
      int q = adj[w].get(i);
      if (!visited[q]) {
        prev[q] = w;
        if (q == t) {
          print(prev, s, t);
          return;
        }
        visited[q] = true;
        queue.add(q);
      }
    }
  }
}

private void print(int[] prev, int s, int t) { // 递归打印s->t的路径
  if (prev[t] != -1 && t != s) {
    print(prev, s, prev[t]);
  }
  System.out.print(t + " ");
}
*/

func (g *Graph) BFS(origin, end int) {
	if origin == end {
		return
	}
	visited := make([]bool, g.VertexNums)
	visited[origin] = true

	queue := []int{origin}
	prev := make([]int, g.VertexNums)
	for i := range prev {
		prev[i] = -1
	}

	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		for i := 0; i < len(g.AdjacencyTable[v]); i++ {
			q := g.AdjacencyTable[v][i]
			if !visited[q] {
				prev[q] = v
				if q == end {
					print(prev, origin, q)
					return
				}
				visited[q] = true
				queue = append(queue, q)
			}
		}
	}
}

func print(prev []int, origin, t int) {
	fmt.Print(prev)
	if prev[origin] != -1 && prev[origin] != prev[t] {
		print(prev, origin, prev[t])
	}

	fmt.Print(fmt.Sprintf("%v ", t))
}
