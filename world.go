package main

World struct {
	nodes 			[][]node 
	height 			int 
	width 			int 
	oneSpaceMoves 	[][]int 
	worldType 		string 
}

func populate(w *world) {

	// species1 intialisation
	species1Lifetime = 10000
	species1Name = 1
	species1Fitness = 0.8
	species1Node := w[0][0] // Should this be a pointer?
	species1 := species{name: species1Name, lifeTime: species1Lifetime, fitness: species1Fitness, node: species1Node}
	species1Node.resident = species1
	go run(species1)

	// species2 initialisation
	species2Lifetime = 5000
	species2Name = 2
	species2Fitness = 0.4
	species2Node := w[4][14] // Should this be a pointer?
	species2 := species{name: species2Name, lifeTime: species2Lifetime, fitness: species2Fitness, node: species2Node}
	species2Node.resident = species2
	go run(species2)
}

func generateOneSpaceMoves() [][]int {
	int numDirections = 9
	int numDimensions = 2
	moves := [numDirections][numDimensions]int

	moves[0,0] = -1 
	moves[0,1] = 1

	moves[1,0] = 0
	moves[1,1] = 1

	moves[2,0] = 1
	moves[2,1] = 1

	moves[3,0] = -1
	moves[3,1] = 0

	moves[4,0] = 0
	moves[4,1] = 0

	moves[5,0] = 1
	moves[5,1] = 0

	moves[6,0] = -1
	moves[6,1] = -1

	moves[7,0] = 0
	moves[7,1] = -1

	moves[8,0] = 1
	moves[8,1] = -1

	return moves
}

func generateNodes(w *world) {
	for i := range len(w.width) {
		for j := range len(w.height) {
			ch := make(chan<- string)
			n = node{horiz: i, vert: j, resident: nil, neighbouringNodes: nil, channelToResident: ch}
			w.nodes[i][j] = n
		}
	}
} 

func assignNeighbours(w *world) {
	for i := range w.height {
		for j := range w.width {
			n :=  // Should this be a pointer?
			select {
				case worldType == "flatWorld":
					getFlatworldNeighbouringNodes()

			}
		}
	}
}
