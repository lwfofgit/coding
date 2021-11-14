package heap_sort

import (
	"fmt"
	"github.com/lwfofgit/coding/src/algorithms/golang/common"
)

// n表示数据的个数，列表a中的数据从下标1到n的位置,列表0不存储数据
// 从下标1开始存储数据，不用0是因为从1开始好计算堆的左右子树位置2*1和2*1+1，可以简化计算
func Sort(data []int, n int) {
	buildHeap(data, n)
	fmt.Println(data)

	for k := n; k > 1; {
		common.Swap(data, 1, k)
		k--
		heapify(data, k, 1)
	}
	fmt.Println(data)
}

func buildHeap(data []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapify(data, n, i)
	}
}

// 从上往下堆化，建立大顶堆，从第一非叶子节点开始(2*i),依次和它的左右子节点对比
func heapify(data []int, n, i int) {
	for {
		maxPos := i
		if i*2 <= n && data[i] < data[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= n && data[maxPos] < data[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}

		common.Swap(data, i, maxPos)
		i = maxPos
	}
}
