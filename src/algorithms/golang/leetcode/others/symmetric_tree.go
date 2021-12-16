package others
import . "github.com/lwfofgit/coding/src/algorithms/golang/leetcode"
/*
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3

思路：
如果一个树的左子树与右子树镜像对称，那么这个树是对称的。
因此，该问题可以转化为：两个树在什么情况下互为镜像？

如果同时满足下面的条件，两个树互为镜像：
它们的两个根结点具有相同的值
每个树的右子树都与另一个树的左子树镜像对称

*/

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

