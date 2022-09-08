package piscine

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	fmt.Printf("BeginningRoot: %v\n", root)
	if root == nil {
		fmt.Printf("NilRoot: %v\n", root)
		return &TreeNode{Data: data}
	} else if data < root.Data {
		fmt.Printf("LeftRootBefRec: %v\n", root)
		root.Left = BTreeInsertData(root.Left, data)
		fmt.Printf("LeftRoot: %v\n", root.Left)
		root.Left.Parent = root
		fmt.Printf("RootLeft: %v\n", root.Left)
	} else {
		fmt.Printf("RightRootBefRec: %v\n", root)
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
		fmt.Printf("RootRight: %v\n", root.Right)
	}
	return root
}
