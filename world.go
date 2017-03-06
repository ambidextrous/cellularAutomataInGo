package main

type world struct {
	nodes 			[][]node 
	height 			int 
	width 			int 
	oneSpaceMoves 	[][]int 
	worldType 		string 
}

func createDefaultWorld() world {
	w := world{nodes: nil, height: 10, width: 30, oneSpaceMoves: nil, worldType: "roundWorld"}
	return w
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
	numDirections := 9
	numDimensions := 2
	moves := [numDirections][numDimensions]int

	moves[0][0] = -1 
	moves[0][1] = 1

	moves[1][0] = 0
	moves[1][1] = 1

	moves[2][0] = 1
	moves[2][1] = 1

	moves[3][0] = -1
	moves[3][1] = 0

	moves[4][0] = 0
	moves[4][1] = 0

	moves[5][0] = 1
	moves[5][1] = 0

	moves[6][0] = -1
	moves[6][1] = -1

	moves[7][0] = 0
	moves[7][1] = -1

	moves[8][0] = 1
	moves[8][1] = -1

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
			&n := w.nodes[i][j] // Pointer
			select {
				case w.worldType == "flatWorld":
					fmt.Println("Making flatWorld")
					neighbours := getFlatWorldNeighbouringNodes(&n,&w)
				case w.worldType == "roundWorld":
					fmt.Println("Making roundWorld")
					neighbours := getRoundWorldNeighbouringNodes(&n,&w)
				default:
					fmt.Println("Making default world (roundWorld)")
					 neighbourss := getRoundWorldNeighbouringNodes(&n,&w)
			}
		}
	}
}

func getRoundWorldNeighbouringNodes(n *node, w *world) []*node {
	numNeighbours := 9
	validNeighbours := make([]*node, numNeighbours)
	for i, oneSpaceMove := range w.oneSpaceMoves {
		hCoordinate := 0
		neighbouringH := n.vert + oneSpaceMove[hCoordinate]
		wCoordinate := 1
		neighbouringW := n.horiz + oneSpaceMove[wCoordinate]
		if neighbouringH == -1 {
			neighbouringH = w.height - 1
		} else if neighbouringH == w.height {
			neighbouringH = 0
		}
		if neighbouringW == -1 {
			neighbouringW = w.width - 1
		} else if neighbouringW == w.width {
			neighbouringW = 0
		}
		neighour := &w.nodes[neighbouringH][neighbouringW]
		validNeighbours[i] = neighour 
	} 
	return validNeighbours
}

func (w world) String() string {
	worldState := ""
	for i := 0 ; i < w.horiz ; i++ {
		for j := 0 ; j < w.vert ; j++ {
			worldState += w.nodes[i][j]
		}
		worldState += "\n"
	} 
	return worldState
}
