package main

import (
	list2 "container/list"
	"fmt"
)

type Node struct {
	Val   interface{}
	Befor *Node
	Next  *Node
}

type List struct {
	Size int
	Head *Node
	Tail *Node
}

func (this *List) InitList() {
	this.Size = 0
	this.Head = nil
	this.Tail = nil
}

//获取第index个节点的值
func (this *List) Get(index int) interface{} {
	if index <= 0 || index > this.Size {
		return -1
	}
	node := this.Head
	for i := 1; i < index; i++ {
		node = node.Next
		if node == nil {
			return -1
		}
	}
	return node.Val
}

//在head之前插入一个节点
func (this *List) AddAtHead(val int) {
	node := new(Node)
	node.Val = val

	if this.Size == 0 {
		this.Head = node
		this.Tail = node
		this.Head.Next = this.Tail
		this.Tail.Befor = this.Head
	} else {
		node1 := this.Head
		node.Next = node1
		node1.Befor = node

		this.Head = node
	}
	this.Size++
}

//在tail之后插入一个节点
func (this *List) AddAtTail(val int) {
	node := new(Node)
	node.Val = val

	if this.Size == 0 {
		this.Head = node
		this.Tail = node
		this.Head.Next = this.Tail
		this.Tail.Befor = this.Head
	} else {
		node1 := this.Tail
		node1.Next = node
		node.Befor = node1

		this.Tail = node
	}
	this.Size++
}

//在第index节点后插入节点
func (this *List) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	node := new(Node)
	node.Val = val
	node1 := this.Head
	for i := 1; i < index; i++ {
		node1 = node1.Next
	}
	node2 := node1.Next

	node.Next = node2
	node.Befor = node1

	node1.Next = node
	node2.Befor = node

	this.Size++
	return
}

//删除第index个节点
func (this *List) DeleteAtIndex(index int) {
	if index <= 0 || index > this.Size {
		return
	}
	if index == this.Size && this.Size == 1 {
		this.Size--
		this.Head = nil
		this.Tail = nil
		return
	}
	if index == this.Size {
		this.Tail = this.Tail.Befor
		this.Tail.Next = nil
		this.Size--
		return
	}
	if index == 1 {
		this.Head = this.Head.Next
		this.Head.Befor = nil
		this.Size--
		return
	}
	node1 := this.Head
	for i := 1; i < index - 1; i++ {
		node1 = node1.Next
	}

	node2 := node1.Next
	node3 := node2.Next

	node1.Next = node3
	node3.Befor = node1
	this.Size--
	return
}

func (this *List) Show() {
	if this.Size == 0 {
		fmt.Println("no data")
	}
	node1 := this.Head
	fmt.Printf("%d: %d\n", 1, node1.Val)
	for i := 1; i < this.Size; i++{
		node1 = node1.Next
		fmt.Printf("%d: %d\n", i + 1, node1.Val)
	}
}

//test list
/*func main() {
	list := new(List)
	list.AddAtTail(4)
	list.Show()
	fmt.Println("-------我是分割线-------")
	list.AddAtHead(2)
	list.Show()
	fmt.Println("-------我是分割线-------")
	list.AddAtHead(1)
	list.Show()
	fmt.Println("-------我是分割线-------")
	list.AddAtTail(5)
	list.Show()
	fmt.Println("-------我是分割线-------")
	list.AddAtIndex(2,3)
	list.Show()
	fmt.Println("-------我是分割线-------")
	list.AddAtIndex(5,3)
	list.Show()
	fmt.Println("-------我是分割线-------")
	list.AddAtIndex(0,3)
	list.Show()
	fmt.Println("-------我是分割线-------")
	list.AddAtIndex(7,10)
	list.Show()
	fmt.Println("-------我是分割线-------")
	list.DeleteAtIndex(3)
	list.Show()
	fmt.Println("-------我是分割线-------")
	fmt.Println(list.Get(1))
}*/

type Queue struct {
	Size int
	Length int
	Head *Node
	Tail *Node
}

func (this *Queue) InitQueue(length int) {
	if length > 0 {
		this.Length = length
	} else {
		fmt.Println("长度必须大于0")
	}
}

func (this *Queue) Front() interface{} {
	if this.Size <= 0 {
		return -1
	}
	return this.Head.Val
}

func (this *Queue) Rear() interface{} {
	if this.Size <= 0 {
		return -1
	}
	return this.Tail.Val
}

func (this *Queue) EnQueue(value interface{}) bool {
	if this.Size >= this.Length {
		return false
	}
	node := new(Node)
	node.Val = value
	if this.Size == 0 {
		this.Head = node
		this.Tail = node
		this.Head.Next = this.Tail
		this.Tail.Befor = this.Head
	} else {
		this.Tail.Next = node
		node.Befor = this.Tail
		this.Tail = node
	}
	this.Size++
	return true
}

func (this *Queue) DeQueue() bool {
	if this.Size == 0 {
		return false
	}
	if this.Size == 1 {
		this.Head = nil
		this.Tail = nil
	} else {
		this.Head = this.Head.Next
	}
	this.Size--
	return true
}

func (this *Queue) QueueShow() {
	if this.Size == 0 {
		return
	}
	node1 := this.Head
	fmt.Printf("%d: %d\n", 1, node1.Val)
	for i := 1; i < this.Size; i++{
		node1 = node1.Next
		fmt.Printf("%d: %d\n", i + 1, node1.Val)
	}
}

func (this *Queue) IsEmpty() bool {
	if this.Size == 0 {
		return true
	}
	return false
}

func (this *Queue) IsFull() bool {
	if this.Size == this.Length {
		return true
	}
	return false
}

/*func main() {
	queue := new(Queue)
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())
	fmt.Println(queue.Front())
	fmt.Println(queue.Rear())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.EnQueue(1))
	queue.InitQueue(3)
	fmt.Println(queue.Front())
	fmt.Println(queue.Rear())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.EnQueue(1))
	fmt.Println(queue.EnQueue(2))
	fmt.Println(queue.EnQueue(3))
	fmt.Println(queue.EnQueue(4))
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.Front())
	fmt.Println(queue.Rear())
	queue.QueueShow()
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.EnQueue(4))
	fmt.Println(queue.EnQueue(4))
	queue.QueueShow()
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())
}
*/
type Stack struct {
	Size int
	Head *Node
	Tail *Node
}

func (this *Stack) Push(x int) {
	node := new(Node)
	node.Val = x
	if this.Size == 0 {
		this.Head = node
		this.Tail = node
		this.Head.Next = this.Tail
		this.Tail.Befor = this.Head
	} else {
		this.Tail.Next = node
		node.Befor = this.Tail
		this.Tail = node
	}
	this.Size++
}

func (this *Stack) Pop() bool {
	if this.Size == 0 {
		return false
	}
	if this.Size == 1 {
		this.Head = nil
		this.Tail = nil
	} else {
		this.Tail = this.Tail.Befor
	}
	this.Size--
	return true
}

func (this *Stack) Top() (interface{}, bool) {
	if this.Size == 0 {
		return -1, false
	}
	return this.Tail.Val, true
}

func (this *Stack) GetMin() (int, bool) {
	if this.Size == 0 {
		return -1, false
	}
	node := this.Head
	min := node.Val.(int)
	for i := 1; i < this.Size; i++ {
		node = node.Next
		if node.Val.(int) < min {
			min  = node.Val.(int)
		}
	}
	return min, true
}

func (this *Stack) StackShow() {
	if this.Size == 0 {
		return
	}
	node1 := this.Head
	fmt.Printf("%d: %d\n", 1, node1.Val)
	for i := 1; i < this.Size; i++{
		node1 = node1.Next
		fmt.Printf("%d: %d\n", i + 1, node1.Val)
	}
}

/*func main() {
	stack := new(Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Push(3)
	stack.StackShow()
	x, ok := stack.Top()
	if ok {
		fmt.Println(x)
	}
	x, ok = stack.GetMin()
	if ok {
		fmt.Println(x)
	}
}*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*func main() {
	root := new(TreeNode)
	root1 := new(TreeNode)
	root2 := new(TreeNode)
	root3 := new(TreeNode)
	root4 := new(TreeNode)
	root5 := new(TreeNode)
	root6 := new(TreeNode)
	root7 := new(TreeNode)

	root.Val = 3
	root1.Val = 9
	root2.Val = 20
	root3.Val = 15
	root4.Val = 7
	root5.Val = 10
	root6.Val = 11
	root7.Val = 12

	root.Left = root1
	root.Right = root2

	root2.Left = root3
	root2.Right = root4

	root4.Left = root5

	root5.Right = root6

	root6.Right = root7

	queue := new(Queue)
	queue.InitQueue(30)

	result := make([][]int, 0)
	data := make([]int, 0)

	ok := queue.EnQueue(root)
	if ok != true {
		return
	}
	var i uint = 0
	var i1 uint = 0
	j := 0
	j1 := 1
	j2 := 0
	j3 := 0
	for queue.IsEmpty() == false {
		node := queue.Front().(*TreeNode)
		ok := queue.DeQueue()
		if ok != true {
			return
		}
		j++
		//fmt.Println(1 << i1, j1, j2, j3)
		if (j1 + j2 + j3) == (1 << i1) {
			j+= j2 + j3
			j3 = (j3 * 2) + (j2 * 2)
			j1 = 0
			j2 = 0
			i1++
		}
		data = append(data, node.Val)
		//fmt.Println(data, 1 << i, j)
		if j == (1 << i) {
			result = append(result, data)
			data = make([]int, 0)
			j = 0
			i++
		}
		if node.Left != nil {
			ok = queue.EnQueue(node.Left)
			if ok != true {
				return
			}
			j1++
		} else {
			j2++
		}
		if node.Right != nil {
			ok = queue.EnQueue(node.Right)
			if ok != true {
				return
			}
			j1++
		} else {
			j2++
		}
	}
	fmt.Println(result)
}*/

func main() {
	root := new(TreeNode)
	root1 := new(TreeNode)
	root2 := new(TreeNode)
	root3 := new(TreeNode)
	root4 := new(TreeNode)
	root5 := new(TreeNode)
	root6 := new(TreeNode)
	root7 := new(TreeNode)

	root.Val = 3
	root1.Val = 9
	root2.Val = 20
	root3.Val = 15
	root4.Val = 7
	root5.Val = 10
	root6.Val = 11
	root7.Val = 12

	root.Left = root1
	root.Right = root2

	root2.Left = root3
	root2.Right = root4

	root4.Left = root5

	root5.Right = root6

	root6.Right = root7

	result := make([][]int, 0)
	data := make([]int, 0)

	list := list2.New()
	list.PushBack(root)

	var i1 uint = 0
	var i uint = 0
	j := 0
	j1 := 1
	j2 := 0
	j3 := 0

	for list.Len() != 0 {
		node1 := list.Front()
		list.Remove(node1)
		node := node1.Value.(*TreeNode)
		j++

		data = append(data, node.Val)

		fmt.Println((1 << i1), j1, j2, j3)
		if (1 << i1) == (j1 + j2 + j3) {
			j = j + j2 + j3
			j3 = (j3 * 2) + (j2 * 2)
			j1 = 0
			j2 = 0
			i1++
		}

		if (1 << i) == j {
			fmt.Println(data, (1 << i), j)
			result = append(result, data)
			data = make([]int, 0)
			j = 0
			i++
		}

		if node.Left != nil {
			list.PushBack(node.Left)
			j1++
		} else {
			j2++
		}

		if node.Right != nil {
			list.PushBack(node.Right)
			j1++
		} else {
			j2++
		}
	}
	fmt.Println(result)
}