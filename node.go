package main

type node struct {
	horiz             int
	vert              int
	resident          species
	neighbouringNodes []node        // Slice needs to be made prior to assignment
	channelToResident chan<- string // Channel needs to be made prior to assignment
}

func murderOccupant(n *node) {
	fmt.Println("Have at you, vile curr!")
	n.channelToResident <- "BANG!"
}

func (n node) String() string {
	emptyNodeSymbol = "-"
	if n.resident == nil {
		return emptyNodeSymbol
	}
	return n.resident
}
