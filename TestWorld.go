package main

import "time"
import "fmt"

// Main struct in which all other simulation elements are located, contains a two-dimensional array of nodes
type world struct {
	nodes         [][]node
	height        int
	width         int
	oneSpaceMoves [][]int
	worldType     string
}

// Returns a string representation of a given world's state
func (w world) String() string {
	worldState := ""
	for i := 0; i < w.width; i++ {
		for j := 0; j < w.height; j++ {
			n := w.nodes[i][j]
			worldState += n.resident.species
			//string nodeState = w.nodes[i][j]
			//worldState += nodeState
		}
		worldState += "\n"
	}
	return worldState
}

// Struct which can be inhabitted by a creature; linked to 0 or 1 creatures by a channel
type node struct {
	horiz             int
	vert              int
	resident          creature
	neighbouringNodes []node        // Slice needs to be made prior to assignment
	channelToResident chan<- string // Channel needs to be made prior to assignment
}

// Returns string representation of a given node
func (n node) String() string {
	emptyNodeSymbol := "-"
	pointerToNode := &n.resident // Pointer
	if pointerToNode == nil {
		return emptyNodeSymbol
	}
	return n.resident.species
}

// Struct representing a node inhabitant, has a species name, a lifetime and a fitness value; linked to a node by a channel
type creature struct {
	species         string
	lifeTime        int
	fitness         float32
	node            *node // Pointer
	channelFromNode <-chan string
}

// Creates a new world struct
func createWorld(height int, width int, worldType string) world {
	w := world{nodes: nil, height: height, width: width, oneSpaceMoves: nil, worldType: worldType}
	w = addNodes(w)
	fmt.Println("w.nodes = ", w.nodes)
	return w
}

// Adds empty nodes to a given  world
func addNodes(w world) world {
	// Creates empty 2D node slices
	w.nodes = make([][]node, w.width)
	for i := 0; i < w.width; i++ {
		w.nodes[i] = make([]node, w.height)
	}
	// Populates slices with empty nodes
	for i := 0; i < w.width; i++ {
		for j := 0; j < w.height; j++ {
			ch := make(chan string)
			n := node{horiz: i, vert: j, neighbouringNodes: nil, channelToResident: ch}
			w.nodes[i][j] = n
		}
	}
	return w
}

// Main programme function
func main() {
	height := 10
	width := 30
	worldType := "roundWorld"
	w := createWorld(height, width, worldType)
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(w)
	}
}
