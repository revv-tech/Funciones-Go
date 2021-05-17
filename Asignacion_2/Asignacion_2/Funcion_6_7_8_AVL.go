package main

//Funcion #2
//Nodo del arbol

type NodeAVL struct {
	//Attributes
	key   int
	left  *NodeAVL
	right *NodeAVL
}

//Representacion de Binary Search Tree
type AVL struct {
	//Attributes
	root *NodeAVL
}

//Funcion #7: Busqueda en AVL
func (avl *AVL) searchAVL(key int) {

}

//Funcion #8: Insercion en AVL
func (avl *AVL) insertNodeAVL(node *NodeAVL, key int) {
	//Crea nuevo node
	newNode := &NodeAVL{key, nil, nil}
	//Verifica si es nulo
	if avl.root == nil {
		//Iguala el nodo
		avl.root = newNode
	} else {
		//Verifica si el key del node es el mismo
		if key == node.key {

		}
		//Verifica si el hijo izq. es mayor
		if key < node.key {
			avl.insertNodeAVL(node.left, key)
		}
		//Verifica si el hijo der. es mayor
		if key > node.key {
			avl.insertNodeAVL(node.right, key)
		}

	}
	return
}

//Balancear el AVL

func (avl *AVL) balancear() *NodeAVL {
	return nil
}

//Diferencia de alturas
func (avl *AVL) difAlturas() int {
	return 0
}

//Rotaciones del arbol
func (avl *AVL) rr_Rotacion(node *NodeAVL) *NodeAVL {
	//Crea un temporal
	tmp := node.right
	node.right = tmp.left
	tmp.left = node
	return tmp

}
func (avl *AVL) ll_Rotacion(node *NodeAVL) *NodeAVL {
	tmp := node.left
	node.right = tmp.right
	tmp.right = node
	return tmp

}
func (avl *AVL) lr_Rotacion(node *NodeAVL) *NodeAVL {

	tmp := node.right
	node.right = avl.rr_Rotacion(tmp)

	return avl.ll_Rotacion(node)

}
func (avl *AVL) rl_Rotacion(node *NodeAVL) *NodeAVL {

	tmp := node.right
	node.right = avl.ll_Rotacion(tmp)
	return avl.rr_Rotacion(node)

}
