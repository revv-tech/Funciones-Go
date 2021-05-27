package main

import (
	"fmt"
	"math"
)

func main() {
	bst := &BST{}
	bst.insertNode(20)
	bst.insertNode(15)
	bst.insertNode(17)
	bst.insertNode(19)
	bst.insertNode(30)
	bst.insertNode(25)
	bst.insertNode(40)
	bst.insertNode(32)
	bst.insertNode(50)
	bst.printBST()
	bst.dsw()
	bst.printBST()

}

//Funcion #2
//Nodo del arbol
type Node struct {
	//Attributes
	key   int
	quant int
	left  *Node
	right *Node
}

//Representacion de Binary Search Tree
type BST struct {
	//Attributes
	root *Node
}

//Funcion #3: Funcion de Buscar
func (bst *BST) search(key int) (int, bool) {
	if bst.root == nil {
		return 1, false
	} else {
		return search_Aux(bst.root, key, 0)
	}
}

//Auxiliar que realiza las comparaciones
func search_Aux(node *Node, key int, recursions int) (int, bool) {

	//Verifica si el nodo es nulo
	if node == nil {
		return recursions, false
	}
	// Verifica si la key es mayor que el la del nodo, si es menor se llama asi mismo con el nodo izquierdo y suma 1 a la cantidad de comparaciones
	if key < node.key {
		return search_Aux(node.left, key, recursions+1)
	}
	// Verifica si la key es menor que el la del nodo, si es mayor se llama asi mismo con el nodo derecho y suma 1 a la cantidad de comparaciones
	if key > node.key {
		return search_Aux(node.right, key, recursions+1)
	}
	return recursions, true

}

//Funcion #4: Insert Function
func (bst *BST) insertNode(key int) int {

	newNode := &Node{key, 1, nil, nil}
	if bst.root == nil {
		bst.root = newNode
		return 1
	} else {
		return insertNodeAux(bst.root, newNode, 0)

	}
}

//Auxiliar que realizar comparaciones
/*
E: dos nodos (uno es la raiz, y el otro el nodo a insertar)
*/
func insertNodeAux(node, newNode *Node, recursions int) int {

	if newNode.key == node.key {
		node.quant = node.quant + 1
		return recursions
	}

	if newNode.key < node.key {

		if node.left == nil {
			recursions++
			node.left = newNode
			return recursions

		} else {

			return insertNodeAux(node.left, newNode, recursions+1)

		}

	} else {

		if node.right == nil {
			recursions++
			node.right = newNode
			return recursions

		} else {

			return insertNodeAux(node.right, newNode, recursions+1)
		}

	}

	return recursions

}

/*Imprimir BST*/
func (bst *BST) printBST() {
	fmt.Println("------------------------------------------------")
	printBSTAux(bst.root, 0)
	fmt.Println("------------------------------------------------")
}
func printBSTAux(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		printBSTAux(n.right, level)
		fmt.Printf(format+"%d\n", n.key)
		printBSTAux(n.left, level)
	}
}

func (bst *BST) dsw() {

	bst.createVine()
	bst.printBST()
	bst.createBalancedBST()

}

func (bst *BST) createVine() {

	tmp := bst.root
	var rounder *Node
	for tmp != nil {
		//bst.printBST()

		if tmp.left != nil {
			left := tmp.left
			rounder = bst.rotateRightDSW(rounder, tmp, left)
			tmp = left
		} else {
			rounder = tmp
			tmp = tmp.right
		}
	}

}

//Funcion #5: DSW
func (bst *BST) createBalancedBST() {

	n := bst.numOfNodes()
	h := int(math.Log2(float64(n + 1)))
	x := h - 1
	m := n - x
	fmt.Println(n, m, x)

	var rounder *Node
	tmp := bst.root
	right := tmp.right

	for x > 0 {

		if right != nil {

			bst.rotateLeftDSW(rounder, tmp, right)
			rounder = right
			tmp = rounder.right
			right = tmp.right

		} else {
			break
		}
		x = x - 1

	}

	for m > 1 {
		var rounder2 *Node
		newTmp := bst.root
		newRight := newTmp.right

		bst.printBST()
		m = m / 2

		if newRight != nil {

			for i := 0; i < m; i++ {

				bst.rotateLeftDSW(rounder2, newTmp, newRight)
				rounder2 = newRight
				newTmp = rounder2.right
				newRight = newTmp.right

			}

		} else {
			break
		}

	}

}

func (bst *BST) rotateRightDSW(rounder, parent, left *Node) *Node {

	if rounder != nil {
		rounder.right = left
	} else {
		bst.root = left
	}
	parent.left = left.right
	left.right = parent

	return rounder
}

func (bst *BST) rotateLeftDSW(rounder, parent, right *Node) {

	if rounder != nil {
		rounder.right = right
	} else {
		bst.root = right
	}
	parent.right = right.left
	right.left = parent

}

func (bst *BST) numOfNodes() int {
	return bst.numOfNodesAux(bst.root)
}
func (bst *BST) numOfNodesAux(node *Node) int {
	if node == nil {
		return 0
	} else {
		return 1 + bst.numOfNodesAux(node.left) + bst.numOfNodesAux(node.right)
	}
}

func (bst *BST) heightBST(node *Node) int {
	if node == nil {
		return -1
	} else {

		leftHeight := bst.heightBST(node.left)
		rightHeight := bst.heightBST(node.right)

		return maxValue(leftHeight, rightHeight) + 1
	}
}

func maxValue(node1, node2 int) int {
	if node1 > node2 {
		return node1
	}
	return node2
}
