package main

import (
	"fmt"
)
type NodeAVL struct {
	key    int
	height int
	quant  int
	left   *NodeAVL
	right  *NodeAVL
}

type TreeAVL struct {
	root                          *NodeAVL
	CantidadComparaciones         int
	CantidadComparacionesBusqueda int
}

/*Imprimir AVL*/
func printAVL(root *NodeAVL) {
	fmt.Println("------------------------------------------------")
	printAVLAux(root, 0)
	fmt.Println("------------------------------------------------")
}
func printAVLAux(n *NodeAVL, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		printAVLAux(n.right, level)
		fmt.Printf(format+"%d\n", n.quant)
		printAVLAux(n.left, level)
	}
}

// Metodo que devuelve el valor de la altura de un nodo
func getHeight(node *NodeAVL) int {

	if node == nil {

		return 0
	} else {

		return node.height
	}
}

// Metodo que crea una nueva instancia de NodeAVL
func newNode(key int) *NodeAVL {

	var newNode *NodeAVL
	newNode = &NodeAVL{key: key, left: nil, right: nil, height: 1, quant: 1}
	return newNode
}

// Metodo que realiza una rotacion hacia la derecha del nodo proporcionado
func rotateRight(y *NodeAVL) *NodeAVL {

	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = max(getHeight(y.left), getHeight(y.right)) + 1
	x.height = max(getHeight(x.left), getHeight(x.right)) + 1

	return x

}

// Metodo que realiza una rotacion hacia la izquierda del nodo proporcionado
func rotateLeft(x *NodeAVL) *NodeAVL {

	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = max(getHeight(x.left), getHeight(x.right)) + 1
	y.height = max(getHeight(y.left), getHeight(y.right)) + 1

	return y
}

// Metodo que calcula el balance del arbol
func getBalance(node *NodeAVL) int {

	if node == nil {
		return 0
	} else {
		return getHeight(node.left) - getHeight(node.right)
	}
}

func (avl *TreeAVL) insertAVL(key int) int {

	if avl.root == nil {
		avl.root = newNode(key)
		globalRecursions = 1
		return globalRecursions
	} else {
		node, _ := insertAVL_Aux(avl.root, key, 0)
		avl.root = node
		return globalRecursions
	}

}

func insertAVL_Aux(node *NodeAVL, key int, recursions int) (*NodeAVL, int) {

	// Se realiza una insercion comun de BST
	if node == nil {
		recursions++
		globalRecursions = recursions
		return newNode(key), recursions
	}

	if key < node.key {
		node.left, _ = insertAVL_Aux(node.left, key, recursions+1)
	} else if key > node.key {
		node.right, _ = insertAVL_Aux(node.right, key, recursions+1)
	} else {
		node.quant = node.quant + 1
		recursions++
		globalRecursions = recursions
		return node, recursions
	}

	//Se actualiza la altura del nodo acenstro
	node.height = 1 + max(getHeight(node.left), getHeight(node.right))

	//Se obtiene el factor de balance de este nodo ancestro para sabre si aun mantiene el balance
	balance := getBalance(node)

	//En caso de estar desbalanceado se tienen 4 posibles casos

	//#1 Caso Izq Izq
	if balance > 1 && key < node.left.key {

		return rotateRight(node), 0

	}
	if balance < -1 && key > node.right.key { //#2 Caso Der Der

		return rotateLeft(node), 0

	}
	if balance > 1 && key > node.left.key { //#3 Caso Izq Der

		node.left = rotateLeft(node.left)
		return rotateRight(node), 0

	}
	if balance < -1 && key < node.right.key { //#4 Caso Der Izq

		node.right = rotateRight(node.right)
		return rotateLeft(node), 0
	}

	//Retornando el node sin ningun tipo de cambio
	return node, recursions
}

//Funcion #3: Funcion de Buscar
func (avl *TreeAVL) searchAVL(root *NodeAVL, key int) (int, bool) {
	if root == nil {
		return 1, false
	} else {
		recursions, bolean := searchAVL_Aux(root, key, 1)
		avl.CantidadComparacionesBusqueda = avl.CantidadComparacionesBusqueda + recursions
		return recursions, bolean
	}
}

//Auxiliar que realiza las comparaciones
func searchAVL_Aux(node *NodeAVL, key int, recursions int) (int, bool) {

	//Verifica si el nodo es nulo
	if node == nil {
		return recursions, false
	}
	// Verifica si la key es mayor que el la del nodo, si es menor se llama asi mismo con el nodo izquierdo y suma 1 a la cantidad de comparaciones
	if key < node.key {
		return searchAVL_Aux(node.left, key, recursions+1)
	}
	// Verifica si la key es menor que el la del nodo, si es mayor se llama asi mismo con el nodo derecho y suma 1 a la cantidad de comparaciones
	if key > node.key {
		return searchAVL_Aux(node.right, key, recursions+1)
	}
	return recursions, true

}

func preOrder(root *NodeAVL) {

	if root != nil {

		fmt.Println(root.key)
		preOrder(root.left)
		preOrder(root.right)
	}
}

func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	}
	return num2
}

func (avl *TreeAVL) numOfNodes() int {
	return avl.numOfNodesAux(avl.root)
}
func (avl *TreeAVL) numOfNodesAux(node *NodeAVL) int {
	if node == nil {
		return 0
	} else {
		return 1 + avl.numOfNodesAux(node.left) + avl.numOfNodesAux(node.right)
	}
}

func profundidad(root *NodeAVL) int {
	if root == nil {
		return 0
	} else {
		left := profundidad(root.left)
		right := profundidad(root.right)

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

