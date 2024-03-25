package binary_tree

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func (t *TreeNode) Search(value int) *TreeNode {
	if t == nil {
		return nil
	}

	if t.value == value {
		return t
	}

	if value > t.value {
		// 大于当前节点，递归右边
		return t.right.Search(value)
	} else {
		// 小于当前节点，递归左边
		return t.left.Search(value)
	}
}

func (t *TreeNode) Insert(value int) {
	if t.value == value {
		return
	}

	if value > t.value {
		if t.right == nil {
			t.right = NewTreeNode(value)
		} else {
			t.right.Insert(value)
		}
	} else {
		if t.left == nil {
			t.left = NewTreeNode(value)
		} else {
			t.left.Insert(value)
		}
	}
}

// DeleteNode 从二叉搜索树中删除值为key的节点，并返回根节点
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.value {
		// 在左子树中删除
		root.left = DeleteNode(root.left, key)
	} else if key > root.value {
		// 在右子树中删除
		root.right = DeleteNode(root.right, key)
	} else {
		// 当前节点就是要删除的节点
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		// 节点有两个子节点，找到右子树的最小节点
		minRight := findMin(root.right)
		root.value = minRight.value                     // 用右子树的最小值替代当前节点的值
		root.right = DeleteNode(root.right, root.value) // 删除那个最小节点
	}
	return root
}

// findMin 查找二叉树中的最小节点
func findMin(node *TreeNode) *TreeNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

// InOrderTraverse 中序遍历
func (t *TreeNode) InOrderTraverse() {
	if t == nil {
		return
	}

	t.left.InOrderTraverse()

	fmt.Println(t.value)

	t.right.InOrderTraverse()

}

// PreOrderTraverse 前序遍历
func (t *TreeNode) PreOrderTraverse() {
	if t == nil {
		return
	}

	fmt.Println(t.value)

	t.left.PreOrderTraverse()

	t.right.PreOrderTraverse()

}

// PostOrderTraverse 后序遍历
func (t *TreeNode) PostOrderTraverse() {
	if t == nil {
		return
	}

	t.left.PostOrderTraverse()
	t.right.PostOrderTraverse()
	fmt.Println(t.value)
}

// LevelOrderTraverse 层序遍历
func (t *TreeNode) LevelOrderTraverse() {
	if t == nil {
		return
	}

	queue := []*TreeNode{t}
	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		queue = queue[1:]
	}
}
