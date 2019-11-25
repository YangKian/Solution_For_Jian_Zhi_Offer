package main

import "Solution_For_Jian_Zhi_Offer/util"

func buildTree(preorder []int, inorder []int) *util.TreeNode {
	size := len(preorder)
	inPos := make(map[int]int, size)
	for k, v := range inorder {
		inPos[v] = k
	}
	return build(&preorder, 0, size-1, 0, &inPos)
}

func build(pre *[]int, preStart, preEnd int, inStart int, inPos *map[int]int) *util.TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := &util.TreeNode{Val: (*pre)[preStart]}
	rootIndex := (*inPos)[(*pre)[preStart]]
	leftLen := rootIndex - inStart

	root.Left = build(pre, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = build(pre, preStart+leftLen+1, preEnd, rootIndex+1, inPos)
	return root
}
