package main

import (
	"fmt"
)

func main() {

	var avl *NodeAVL
	avl := &NodeAVL{}

	avl = insertAVL(avl, 10)
	avl = insertAVL(avl, 20)
	avl = insertAVL(avl, 30)
	avl = insertAVL(avl, 40)
	avl = insertAVL(avl, 50)
	avl = insertAVL(avl, 25)
	preOrder(root)
}

type NodeAVL struct {
	key    int
	height int
	left   *NodeAVL
	right  *NodeAVL
}

// Metodo que devuelve el valor de la altura de un nodo
func (nodeAVL *NodeAVL) getHeight(node *NodeAVL) int {

	if node == nil {

		return 0
	} else {

		return node.height
	}
}

// Metodo que crea una nueva instancia de NodeAVL
func (node *NodeAVL) newNode(key int) *NodeAVL {

	var newNode *NodeAVL
	newNode := new(Node)
	newNode.key = key
	newNode.left = nil
	newNode.right = nil
	newNode.height = 1

	return newNode
}

// Metodo que realiza una rotacion hacia la derecha del nodo proporcionado
func (node *NodeAVL) rotateRight(y *NodeAVL) *NodeAVL {

	var x *NodeAVL
	x := y.left
	var T2 *NodeAVL
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = max(NodeAVL.getHeight(y.left), getHeight(y.right)) + 1
	x.height = max(getHeight(x.left), getHeight(x.right)) + 1

	return x

}

// Metodo que realiza una rotacion hacia la izquierda del nodo proporcionado
func (node *NodeAVL) rotateLeft(y *NodeAVL) *NodeAVL {

	var y *NodeAVL
	y := x.right
	var T2 *NodeAVL
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = max(getHeight(x.left), getHeight(x.right)) + 1
	y.height = max(getHeight(y.left), getHeight(y.right)) + 1

	return y
}

// Metodo que calcula el balance del arbol
func (nodeAVL *NodeAVL) getBalance(node *NodeAVL) int {

	if node == nil {
		return 0
	} else {
		return getHeight(node.left) - getHeight(node.right)
	}
}

//
func (nodeAVL *NodeAVL) insertAVL(node *NodeAVL, key int) *NodeAVL {

	// Se realiza una insercion comun de BST
	if node == nil {
		return newNode(key)
	}

	if key < node.key {
		node.left = insertAVL(node.left, key)
	} else if key > node.key {
		node.right = insertAVL(node.right, key)
	} else {
		return node
	}

	//Se actualiza la altura del nodo acenstro
	node.height = 1 + max(getHeight(node.left), getHeight(node.right))

	//Se obtiene el factor de balance de este nodo ancestro para sabre si aun mantiene el balance
	balance := getBalance(node)

	//En caso de estar desbalanceado se tienen 4 posibles casos

	//#1 Caso Izq Izq
	if balance > 1 && key < node.left.key {

		return NodeAVL.rotateRight(node)
	} else if balance < -1 && key > node.right.key { //#2 Caso Der Der

		return leftRotate(node)
	} else if balance > 1 && key > node.left.key { //#3 Caso Izq Der

		node.left = leftRotate(node)
		return rightRotate(node)
	} else if balance < -1 && key < node.right.key { //#4 Caso Der Izq

		node.right = NodeAVL.rightRotate(node.right)
		return leftRotate(node)
	}

	//Retornando el node sin ningun tipo de cambio
	return node
}

func (node *NodeAVL) preOrder(root *NodeAVL) {

	if root == nil {

		fmt.Println(root.key)
		preOrder(root.left)
		preOrder(root.right)
	}
}

func max(num1, num2 int) int {

	if num1 > num2 {
		return num2
	}
	return node2
}
