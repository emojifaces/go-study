package main

type TreeNode struct {
	value  int
	height int
	left   *TreeNode
	right  *TreeNode
}

// NewTreeNode 创建新的AVL树节点
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		value:  value,
		height: 1,
	}
}

// 获取节点的高度
func (t *TreeNode) getHeight() int {
	if t == nil {
		return 0
	}

	return t.height
}

// 计算节点的平衡因子
func (t *TreeNode) getBalance() int {
	if t == nil {
		return 0
	}

	return t.left.getHeight() - t.right.getHeight()
}

func (t *TreeNode) updateHeight() {
	leftHeight := t.left.getHeight()
	rightHeight := t.right.getHeight()
	t.height = max(leftHeight, rightHeight) + 1
}

// 左旋
func (t *TreeNode) rotateLeft() *TreeNode {

	if t == nil {
		return nil
	}

	// t是左旋之前的根节点

	newRoot := t.right // 左旋之后的根节点

	t.right = newRoot.left
	newRoot.left = t

	t.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// 右旋
func (t *TreeNode) rotateRight() *TreeNode {

	if t == nil {
		return nil
	}

	newRoot := t.left

	t.left = newRoot.right
	newRoot.right = t

	t.updateHeight()
	newRoot.updateHeight()
	return newRoot
}

func (t *TreeNode) Insert(value int) *TreeNode {
	if t == nil {
		return NewTreeNode(value)
	}

	if value < t.value {
		t.left = t.left.Insert(value)
	} else if value > t.value {
		t.right = t.right.Insert(value)
	} else { // 不允许有相同值的节点
		return t
	}

	// 更新高度
	t.updateHeight()

	// 获取平衡因子，检查是否失衡
	balance := t.getBalance()

	// 左左情况
	if balance > 1 && value < t.left.value {
		return t.rotateRight()
	}

	// 右右情况
	if balance < -1 && value > t.right.value {
		return t.rotateLeft()
	}

	// 左右情况
	if balance > 1 && value > t.left.value {
		t.left = t.left.rotateLeft()
		return t.rotateRight()
	}

	// 右左情况
	if balance < -1 && value < t.right.value {
		t.right = t.right.rotateRight()
		return t.rotateLeft()
	}

	// 已平衡，返回节点
	return t
}
