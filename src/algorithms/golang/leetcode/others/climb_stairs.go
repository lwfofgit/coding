package others
/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
*/

func climbStairs(n int) int {
	if n < 0 {
		return 0
	} else if n == 1 {
		return n
	} else if n == 2 {
		return n
	}else {
		return climbStairs(n-1) + climbStairs(n-2)
	}
}

// 滚动数组 f(3) = f(2) + f(1)
func climbStairsV2(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
