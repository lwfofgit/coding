package string_match

import "fmt"

/*
BF 算法中的 BF 是 Brute Force 的缩写，中文叫作暴力匹配算法，也叫朴素匹配算法。
从名字可以看出，这种算法的字符串匹配方式很“暴力”，当然也就会比较简单、好懂，但相应的性能也不高。
在开始讲解这个算法之前，我先定义两个概念，方便我后面讲解。

它们分别是主串和模式串。这俩概念很好理解，我举个例子你就懂了。
比方说，我们在字符串 A 中查找字符串 B，那字符串 A 就是主串，字符串 B 就是模式串。
我们把主串的长度记作 n，模式串的长度记作 m。因为我们是在主串中查找模式串，所以 n>m。

作为最简单、最暴力的字符串匹配算法，BF 算法的思想可以用一句话来概括，那就是，
我们在主串中，检查起始位置分别是 0、1、2....n-m 且长度为 m 的 n-m+1 个子串，看有没有跟模式串匹配的。

BF 算法的时间复杂度很高，是 O(n*m)，但在实际的开发中，它却是一个比较常用的字符串匹配算法。
*/

func BK(parent, child string) int {
	if len(parent) <= len(child) {
		return -1
	}

	for i:=0; i<len(parent)-len(child); i++{
		pc := parent[i: i+len(child)]
		fmt.Println(pc)
		if pc == child {
			return i
		}
	}
	return -1
}
