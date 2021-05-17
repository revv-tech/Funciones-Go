package main

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
		return 0, false
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
		return 0
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
		return recursions + 1
	}

	if newNode.key < node.key {

		if node.left == nil {

			node.left = newNode
			return recursions + 1

		} else {
			insertNodeAux(node.left, newNode, recursions+1)

		}
	} else {

		if node.right == nil {

			node.right = newNode
			return recursions + 1

		} else {
			insertNodeAux(node.right, newNode, recursions+1)
		}

	}
	return 0
}
