package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ArrToTreeNode(arr []interface{}) *TreeNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	var nodeList []*TreeNode
	for _, value := range arr {
		if value != nil {
			nodeList = append(nodeList, &TreeNode{Val: value.(int)})
		}else{
			nodeList = append(nodeList, nil)
		}
	}

	for i := 0; i < len(arr)/2; i++ {
		if nodeList[2*i+1] != nil {
			//左节点
			nodeList[i].Left = nodeList[2*i+1]
		}
		if nodeList[2*i+2] != nil {
			//右节点
			nodeList[i].Right = nodeList[2*i+2]
		}
	}

	return nodeList[0]
}
