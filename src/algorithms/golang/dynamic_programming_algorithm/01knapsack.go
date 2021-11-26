package dynamic_programming_algorithm

/*
对于一组不同重量、不可分割的物品，我们需要选择一些装入背包，
在满足背包最大重量限制的前提下，背包中物品总重量的最大值是多少呢？

假设：
  背包的最大承载重量是 9。我们有 5 个不同的物品，每个物品的重量分别是 2，2，4，6，3。
*/

/*
weight:物品重量，n:物品个数，w:背包可承载重量
public int knapsack(int[] weight, int n, int w) {
  boolean[][] states = new boolean[n][w+1]; // 默认值false
  states[0][0] = true;  // 第一行的数据要特殊处理，可以利用哨兵优化
  if (weight[0] <= w) {
    states[0][weight[0]] = true;
  }
  for (int i = 1; i < n; ++i) { // 动态规划状态转移
    for (int j = 0; j <= w; ++j) {// 不把第i个物品放入背包
      if (states[i-1][j] == true) states[i][j] = states[i-1][j];
    }
    for (int j = 0; j <= w-weight[i]; ++j) {//把第i个物品放入背包
      if (states[i-1][j]==true) states[i][j+weight[i]] = true;
    }
  }
  for (int i = w; i >= 0; --i) { // 输出结果
    if (states[n-1][i] == true) return i;
  }
  return 0;
}
*/

// weight:物品重量，n:物品个数，w:背包可承载重量
func Knapsack(weight []int, n, w int) int {
	states := make([][]bool, n)
	for i := range states {
		states[i] = make([]bool, w+1)
	}

	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}

		for j := 0; j <= w-weight[i]; j++ {
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- {
		if states[n-1][i] == true {
			return i
		}
	}
	return 0
}
