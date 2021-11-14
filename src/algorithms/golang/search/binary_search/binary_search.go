package binary_search

// 注意点：
// 1 low<=high, 不是low < high
// 2 查找时间复杂度 logn, 最快的排序算法时间复杂度是：nlogn
// 3 low = mid+1/mid -1; 如果直接写成low = mid或high； 会发生死循环
func Search(a []int, data int) int {
	n := len(a)
	low := 0
	high := n - 1

	for low <= high {
		mid := low + (high-low)/2
		if a[mid] == data {
			return mid
		} else if a[mid] < data {
			low = mid + 1
		} else {
			low = high - 1
		}
	}
	return -1
}

func RecursionSearch(a []int, data int) int {
	return recursionSearch(a, 0, len(a)-1, data)
}

func recursionSearch(a []int, low, high, data int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)>>1
	if a[mid] == data {
		return mid
	} else if a[mid] < data {
		return recursionSearch(a, low+1, high, data)
	} else {
		return recursionSearch(a, low, high-1, data)
	}
}
