package main

//Funcion #2
//Nodo del arbol
type Node struct {
	//Attributes
	key    int
	height int
	left   *Node
	right  *Node
}

//Representacion de Binary Search Tree
type AVL struct {
	//Attributes
	root *Node
}
