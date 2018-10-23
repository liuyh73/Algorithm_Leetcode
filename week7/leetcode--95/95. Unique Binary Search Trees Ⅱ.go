/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return findTrees(1, n)
}

func findTrees(s int,t int) []*TreeNode {
    list := make([]*TreeNode, 0)
    if s>t {
        list = append(list, nil)
        return list
    }
    if s==t {
        list = append(list, &TreeNode{s, nil, nil})
        return list
    }
    for i:=s; i<=t; i++ {
        leftNodes := findTrees(s, i-1)
        rightNodes := findTrees(i+1, t)
        
        for _, leftNode := range leftNodes {
            for _, rightNode := range rightNodes {
                root := &TreeNode{i, leftNode, rightNode}
                list = append(list, root)
            }
        }
    }
    return list
}