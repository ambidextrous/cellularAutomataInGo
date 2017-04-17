package main

//import "time"
import "fmt"

// Main struct in which all other simulation elements are located, contains a two-dimensional array of nodes
type world struct {
	nodes         [][]node
	width         int
	height        int
	oneSpaceMoves [][]int
	worldType     string
}

// Returns a string representation of a given world's state
func (w world) printWorld() {
	space := " "
	worldState := ""
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			n := w.nodes[i][j]
			residentPointer := &n.resident
			if residentPointer.species == "" {
				worldState += "-"
			} else {
				worldState += n.resident.species
			}
			worldState += space
		}
		worldState += "\n\n"
	}
	fmt.Println(worldState)
}

// Struct which can be inhabitted by a creature; linked to 0 or 1 creatures by a channel
type node struct {
	horiz             int
	vert              int
	resident          creature
	neighbouringNodes []*node     // Slice needs to be made prior to assignment
	channelToResident chan string // Channel needs to be made prior to assignment
}

// Struct representing a node inhabitant, has a species name, a lifetime and a fitness value; linked to a node by a channel
type creature struct {
	species         string
	lifeTime        int
	fitness         float32
	node            *node // Pointer
	channelFromNode chan string
}

// Creates a new world struct
func createWorld(width int, height int, worldType string) world {
	w := world{nodes: nil, width: width, height: height, oneSpaceMoves: nil, worldType: worldType}
	wPointer := &w
	addNodes(wPointer) // Running function on pointer, no return value
	//addAntiCreatures(wPointer)     // Running function on pointer, no return value
	addNeighbouringNodes(wPointer) // Running function on pointer, no return value
	return w
}

// Add roundWorld - where the edges of the world "wrap around" - neighbouring nodes to each node in a given world
func addRoundworldNeighbours(w *world) { // Recives a pointer
	// Create array of one-space moves
	possibleMoves := []int{-1, 0, 1}
	// Iterate through all nodes in world
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			n := &w.nodes[i][j]                  // Pointer
			neighboursCoords := make([][]int, 9) // Slice
			counter := 0
			// Calculate neighbouring node coordinates
			for k := 0; k < len(possibleMoves); k++ {
				for l := 0; l < len(possibleMoves); l++ {
					heightCoord := i + possibleMoves[k]
					if heightCoord == -1 {
						heightCoord += w.height
					}
					widthCoord := j + possibleMoves[l]
					if widthCoord == -1 {
						widthCoord += w.width
					}
					neighbourCoords := []int{(heightCoord % w.height), (widthCoord % w.width)}
					neighboursCoords[counter] = neighbourCoords
					counter++
				}
			}
			// Add neighbours to node
			neighbours := make([]*node, 9) // Pointer
			for m := 0; m < len(neighbours); m++ {
				neighbour := &w.nodes[neighboursCoords[m][0]][neighboursCoords[m][1]]
				neighbours[m] = neighbour
			}
			n.neighbouringNodes = neighbours
		}
	}
	// Test print
	//for i := 0; i < w.height; i++ {
	//	for j := 0; j < w.width; j++ {
	//		fmt.Println(w.nodes[i][j].horiz, w.nodes[i][j].vert)
	//		for k := 0; k < len(w.nodes[i][j].neighbouringNodes); k++ {
	//			fmt.Println(w.nodes[i][j].neighbouringNodes[k].horiz, w.nodes[i][j].neighbouringNodes[k].vert)
	//		}
	//		fmt.Println("")
	//	}
	//}
}

// Calls appropriate neighbouring node assignment function depending on worldType
func addNeighbouringNodes(w *world) {
	switch worldType := w.worldType; worldType {
	default:
		addRoundworldNeighbours(w)
	}
}

// Adds an anti-creature to each node in a given world
//func addAntiCreatures(w *world) {
//	for i := 0; i < w.height; i++ {
//		for j := 0; j < w.width; j++ {
//			nPointer := &w.nodes[i][j]
//			c := creature{species: "-", lifeTime: 1000000000000, fitness: 0.0, node: nPointer, channelFromNode: nPointer.channelToResident}
//			nPointer.resident = c
//		}
//	}
//}

// Adds empty nodes to a given  world
func addNodes(w *world) {
	// Creates empty 2D node slices
	w.nodes = make([][]node, w.height)
	for i := 0; i < w.height; i++ {
		w.nodes[i] = make([]node, w.width)
	}
	// Populates slices with empty nodes
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			ch := make(chan string)
			n := node{horiz: i, vert: j, neighbouringNodes: nil, channelToResident: ch}
			w.nodes[i][j] = n
		}
	}
}

// Main programme function
func main() {
	width := 30
	height := 10
	worldType := "roundWorld"
	w := createWorld(width, height, worldType)

	w.printWorld()
	//fmt.Println(w)
	//for {
	//	time.Sleep(500 * time.Millisecond)
	//	fmt.Println(w)
	//}
}
