package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 自定义工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value} // 返回的是局部变量的地址
}

// 函数名前面括号中是接收者(调用者)，node参数也是值传递
// 若想改为引用传递，需要变成指针 node *Node
func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}
