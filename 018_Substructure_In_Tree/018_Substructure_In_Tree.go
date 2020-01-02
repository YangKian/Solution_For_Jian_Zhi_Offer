package main

import ."Solution_For_Jian_Zhi_Offer/util"

func HasSubtree(s, t *TreeNode) bool {
	var res bool
	if s != nil && t != nil {
		if s.Val == t.Val {
			res = DoesTree1HasTree2(s, t)
		}
	}

	if !res {
		res = HasSubtree(s.Left, t) || HasSubtree(s.Right, t)
	}

	return res
}

func DoesTree1HasTree2(s, t *TreeNode) bool {
	//注意递归的终止条件：到达任意一棵树的叶子节点
	if t == nil {
		return true
	}

	if s == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return DoesTree1HasTree2(s.Left, t.Left) && DoesTree1HasTree2(s.Right, t.Right)
}