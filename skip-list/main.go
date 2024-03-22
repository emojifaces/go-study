package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 跳表

const (
	P = 0.5 // 概率
)

// Node 跳表节点
type Node struct {
	// 节点存储的值
	value int

	// 节点跨越的层数
	// 如果一个节点的level值为3，这意味着这个节点存在于跳表的第0层、第1层和第2层。
	// 实际上，它在这三个层级中都有链接（即next数组有3个有效元素）。
	level int

	// 数组中的每个元素都是指向同一层级中下一个节点的指针
	// next[i]：对于节点A的next数组中的第i个元素，它表示在跳表的第i层中，A节点后面紧跟着的节点。如果next[i]为nil，这表示在第i层中，A节点后面没有节点。
	// eg:
	// 假设我们有一个跳表节点A，它出现在3个层级中，那么A的next字段将是一个长度为3的数组（假设我们从0开始计数层级）：
	// A.next[0]指向同一底层中A的下一个节点。
	// A.next[1]在第1层中指向A之后的节点。
	// A.next[2]在第2层中指向A之后的节点。
	next []*Node
}

// SkipList 跳表
type SkipList struct {
	// 头节点
	head *Node
	// 当前跳表层数
	level int
	// 跳表长度（底层链表的长度）
	length int
	// 最大层数
	maxLevel int
}

// NewNode 创建跳表节点
func NewNode(value int, level int) *Node {
	return &Node{
		value: value,
		level: level,
		next:  make([]*Node, level),
	}
}

// NewSkipList 新建跳表
func NewSkipList(maxLevel int) *SkipList {
	return &SkipList{
		head:     NewNode(0, maxLevel), // 初始化头节点 值为0 层数为最大层
		level:    1,                    // 初始层数为1
		length:   0,                    // 初始长度为0
		maxLevel: maxLevel,             // 记录最大层数
	}
}

// 随机生成新节点的层级
func (s *SkipList) randomLevel() int {
	level := 1 // 默认1层

	for rand.Float64() < P && level < s.maxLevel {
		level++
	}

	return level
}

// Insert 插入值
func (s *SkipList) Insert(value int) {

	update := make([]*Node, s.maxLevel) // 用于记录每层需要更新的节点
	current := s.head                   // 从头节点开始

	// 从最高层开始，找到每层中第一个大于等于插入值的节点的前驱
	for i := s.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	// 获取新节点的层级
	newNodeLevel := s.randomLevel()

	// 如果新节点层数超过当前跳表层数，需要更新跳表的层数
	if newNodeLevel > s.level {
		for i := s.level; i < newNodeLevel; i++ {
			update[i] = s.head
		}
		s.level = newNodeLevel
	}

	// 创建新节点
	newNode := NewNode(value, newNodeLevel)

	for i := 0; i < newNodeLevel; i++ {
		// 新节点的next指向前一个节点的next
		newNode.next[i] = update[i].next[i]
		// 前一个节点的next指向新节点
		update[i].next[i] = newNode
	}

	s.length++
}

// Search 查找值
func (s *SkipList) Search(value int) bool {

	current := s.head
	for i := s.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
	}

	// 移动到下一层的起点，准备搜索下一层
	current = current.next[0]

	// 如果当前节点的值等于目标值，则表示找到了目标值
	return current != nil && current.value == value
}

// Delete 删除值
func (s *SkipList) Delete(value int) {

	update := make([]*Node, s.maxLevel)
	current := s.head

	for i := s.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	current = current.next[0]
	if current != nil && current.value == value {
		for i := 0; i < s.level; i++ {
			if update[i].next[i] != current {
				break
			}
			update[i].next[i] = current.next[i]
		}

		for s.level > 1 && s.head.next[s.level-1] == nil {
			s.level--
		}

		s.length--
	}
}

// 初始化随机种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	sl := NewSkipList(10)

	sl.Insert(4)
	sl.Insert(8)
	sl.Insert(5)
	sl.Insert(9)
	sl.Insert(2)

	fmt.Printf("max level: %d level: %d length:%d \n", sl.maxLevel, sl.level, sl.length) // max level: 3 level: 2 length:5

	fmt.Println(sl.Search(10)) // false
	fmt.Println(sl.Search(5))  // true

	sl.Delete(5)
	sl.Delete(4)

	fmt.Printf("max level: %d level: %d length:%d \n", sl.maxLevel, sl.level, sl.length) // max level: 3 level: 2 length:3

	fmt.Println(sl.Search(5))  // false
	fmt.Println(sl.Search(10)) // false
	fmt.Println(sl.Search(2))  // true
}
