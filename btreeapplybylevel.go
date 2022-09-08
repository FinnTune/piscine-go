package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	treeLvl := BTreeLevelCount(root)
	for i := 1; i <= treeLvl; i++ {
		TreeApplyByLevelRecursion(root, f, i)
	}
}

func TreeApplyByLevelRecursion(root *TreeNode, f func(...interface{}) (int, error), level int) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		TreeApplyByLevelRecursion(root.Left, f, level-1)
		TreeApplyByLevelRecursion(root.Right, f, level-1)
	}
}
