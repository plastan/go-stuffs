package main

import "fmt"



type Node struct{
key int
left *Node
right *Node
height int
}


func max(a,b int) int {
	if a > b { 
		return a
	}
	return b
}



// generate new node

func newNode(key int) *Node {
	node := &Node{key:key}
	node.left = nil
	node.right = nil
	node.height = 1
	return node
}




func height(N *Node) int {
	if N==nil {
		return 0
	}
	return N.height 
}

// calculates the balance factor of the node
func getBalanceFactor (N * Node) int {
	if N == nil {
		return 0
	}
	return height(N.left) - height(N.right)
}


// ROTATIONS

// left rotate
func leftRotate( x *Node) *Node{
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2
	
	x.height = max(height(x.left),height(x.right)) + 1;
	y.height = max(height(y.left),height(y.right)) + 1;
		
	return y
}


//right rotate
func rightRotate( y *Node) *Node{
	x := y.left
	t2 := x.right

	x.right = y
	y.left = t2

	x.height = max(height(x.left),height(x.right)) + 1;
	y.height = max(height(y.left),height(y.right)) + 1;
		
	return y
}
// insertion

func insertNode(node *Node, key int) *Node {
	if node == nil {
		return newNode(key)
	}

	if key < node.key {
		node.left = insertNode(node.left,key)
	}else if key > node.key {
		node.right = insertNode(node.right,key)
	}else {
		return node
	}


	node.height = 1 + max(height(node.left), height(node.right))
	balanceFactor := getBalanceFactor(node)

	if balanceFactor > 1 {
		if key < node.left.key {
			return rightRotate(node)
		}else if key > node.left.key { 
			node.left = leftRotate(node.left)
			return rightRotate(node)
		}
	}

	if balanceFactor < -1 {
		if key > node.right.key {
			return leftRotate(node)
		} else if key < node.right.key { 
			node.right = rightRotate(node.right)
			return leftRotate(node)
		}
	} 

	return node
}
// node with minimm value

func nodeWithMinimumValue(node * Node) *Node {
	current := node
	for current.left != nil{
		current = current.left
	}
	return current
}

// deletion



//search


func printTree(root *Node, indent string, last bool){
	if root != nil{
		fmt.Print(indent)
		if last {
			fmt.Print("R ----")
			indent += "	"
		} else{
			fmt.Print("L ----")
			indent += "|    "
		}
		fmt.Println(root.key)
		printTree(root.left,indent, false)
		printTree(root.right, indent,true)
	}
}




// Main

func main(){
	root := insertNode(nil,22);
	root = insertNode(root,1);
	root = insertNode(root,21);
	root = insertNode(root,12);
	root = insertNode(root,32);
	root = insertNode(root,123);
	root = insertNode(root,43);
	root = insertNode(root,4);
	root = insertNode(root,53);

	printTree(root,"",true)

}

