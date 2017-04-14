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
			worldState += *w.nodes[i][j]
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
	if n.resident == nil {
		return emptyNodeSymbol
	}
	return n.resident
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
