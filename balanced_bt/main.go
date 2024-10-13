package main

import (
	"fmt"
	"strings"
)

type Node struct {
	key    int
	left   *Node
	right  *Node
	height int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// generate new node

func newNode(key int) *Node {
	node := &Node{key: key}
	node.left = nil
	node.right = nil
	node.height = 1
	return node
}

func height(N *Node) int {
	if N == nil {
		return 0
	}
	return N.height
}

// calculates the balance factor of the node
func getBalanceFactor(N *Node) int {
	if N == nil {
		return 0
	}
	return height(N.left) - height(N.right)
}

// ROTATIONS

// left rotate
func leftRotate(x *Node) *Node {
	y := x.right
	t2 := y.left

	y.left = x
	x.right = t2

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

// right rotate
func rightRotate(y *Node) *Node {
	x := y.left
	t2 := x.right

	x.right = y
	y.left = t2

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

// insertion

func insertNode(node *Node, key int) *Node {
	if node == nil {
		return newNode(key)
	}

	if key < node.key {
		node.left = insertNode(node.left, key)
	} else if key > node.key {
		node.right = insertNode(node.right, key)
	} else {
		return node
	}

	node.height = 1 + max(height(node.left), height(node.right))
	balanceFactor := getBalanceFactor(node)

	if balanceFactor > 1 {
		if key < node.left.key {
			return rightRotate(node)
		} else if key > node.left.key {
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

func nodeWithMinimumValue(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

// deletion

// search
func prettyPrint(root *Node, indent int) {

	if root != nil {
		if root.right != nil {
			prettyPrint(root.right, indent+4)
		}
		if indent > 0 {
			fmt.Print(strings.Repeat(" ", indent))
		}
		if root.right != nil {
			fmt.Printf(" /\n")
			fmt.Print(strings.Repeat(" ", indent))
		}
		fmt.Print(root.key, "\n")

		if root.left != nil {
			fmt.Print(strings.Repeat(" ", indent))
			fmt.Printf("\\\n")
			prettyPrint(root.left, indent+4)
		}
	}
}

func printTree(root *Node, indent string, last bool) {
	if root != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R ----")
			indent += "	"
		} else {
			fmt.Print("L ----")
			indent += "|    "
		}
		fmt.Println(root.key)
		printTree(root.left, indent, false)
		printTree(root.right, indent, true)
	}
}

//deletion

func remove(root *Node, key int) *Node {
	// STEP 1: PERFORM STANDARD BST DELETE
	// STEP 2: UPDATE HEIGHT OF THE CURRENT NODE
	// STEP 3: GET THE BALANCE FACTOR OF THIS TO CHECK WEATHER THE NODE IS BALANCED OR NOT
	// STEP 4: BALANCE THE TREE

	// STEP 1: PERFORM STANDARD BST DELETE
	if root.key > key {
		remove(root.left, key)
	} else if root.key < key {
		remove(root.right, key)
	} else {

		// node found
		//define temp
		var temp *Node
		// node with no child
		if root.left == nil || root.right == nil {

			if root.left != nil {
				temp = root.left
			} else {
				temp = root.right
			}

			// no child case

			if temp == nil {
				root = temp
			}

		} else {
			// 2 child present
			temp := nodeWithMinimumValue(root.right)

			root.key = temp.key
			root.right = remove(root.right, temp.key)
		}
	}

	if root == nil {
		return root
	}

	// STEP 2: UPDATE HEIGHT OF THE CURRENT NODE

	root.height = 1 + max(height(root.left), height(root.right))

	// STEP 3: GET THE BALANCE FACTOR OF THIS TO CHECK WEATHER THE NODE IS BALANCED OR NOT

	balance := getBalanceFactor(root)

	// left left
	if balance > 1 && getBalanceFactor(root.left) >= 0 {
		return rightRotate(root)
	}

	// left right

	if balance > 1 && getBalanceFactor(root.left) < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	// right right

	if balance < -1 && getBalanceFactor(root.right) >= 0 {
		return leftRotate(root)
	}

	// right left

	if balance < -1 && getBalanceFactor(root.right) < 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root

}

// travers
func preorder(n *Node) {
	if n != nil {
		fmt.Print(n.key, " ")
		preorder(n.left)
		preorder(n.right)
	}
}

// travers
func inorder(n *Node) {
	if n != nil {
		preorder(n.left)
		fmt.Print(n.key, " ")
		preorder(n.right)
	}
}

func postorder(n *Node) {
	if n != nil {
		preorder(n.left)
		preorder(n.right)
		fmt.Print(n.key, " ")
	}
}

// Main

func main() {
	root := insertNode(nil, 1)
	root = insertNode(root, 2)
	root = insertNode(root, 3)
	root = insertNode(root, 4)
	root = insertNode(root, 5)
	root = insertNode(root, 6)
	root = insertNode(root, 7)
	root = insertNode(root, 8)

	//	printTree(root, "", true)
	//	preorder(root)
	//	fmt.Println()
	//	inorder(root)
	fmt.Println()
	//	postorder(root)
	//	fmt.Println()
	//	remove(root, 2)
	prettyPrint(root, 0)
	// printTree(root, "", true)
}
