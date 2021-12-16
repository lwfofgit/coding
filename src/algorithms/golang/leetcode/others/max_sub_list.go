package others
/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。

题解：
	假设nums 数组的长度是 n，下标从 00到 n-1。我们用 f(i)代表以第 i 个数结尾的「连续子数组的最大和」，
那么很显然我们要求的答案就是：
	0≤i≤n−1：max{f(i)}

递归方程：
f(i)=max{f(i−1)+nums[i],nums[i]}
*/

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i< len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i]  > max {
			max = nums[i]
		}
	}
	return max
}