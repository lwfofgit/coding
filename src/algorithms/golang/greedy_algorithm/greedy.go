package greedy_algorithm

import (
	"strconv"
)

/*
在一个非负整数 a 中，我们希望从中移除 k 个数字，让剩下的数字值最小，如何选择移除哪 k 个数字呢？
解题思路：
1.这里假设非负整数为：7658365398973，它的数字个数是固定的，有一个局限值；
2.从中移除k位后，假设k是5，期待剩下的值最小，有一个期望值，符合贪心算法思路；
3.可以从高位开始循环，和比它低一位值做对比，如果高位比低位大，就移除它，若高位小则向右移动一位继续比较两个数字；
  若遍历到最末位，都不满足上述条件，则末位值是最大的，移除末位；重复上述操作k
*/

func MinValue(value, times int) int {
	if value < 0 {
		return -1
	}
	if value < times {
		return 0
	}

	vs := strconv.Itoa(value)
	for i := 0; i < times; i++ {
		for j := 0; j < len(vs)-1; j++ {
			if vs[j] > vs[j+1] {
				vs = vs[:j] + vs[j+1:]
				break
			}

			if j == len(vs) - 2 {
				vs = vs[:len(vs)-1]
			}
		}
	}

	vi, _ := strconv.Atoi(vs)
	return vi
}

func TestStr(str string, d int) string  {
	return str[:d] + str[d+1:]
}


// 获取一个正整数的个数
func GetNums(value, nums int) int {
	if value < 0 {
		return -1
	}

	if value < 10 {
		return nums + 1
	}

	value = value / 10
	nums++
	return GetNums(value, nums)
}
