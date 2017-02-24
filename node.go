package main

type node stuct {
	int horiz
	int vert
	species resident
	[]node neighbouringNodes // Slice needs to be made prior to assignment
	chan<- string channelToResident // Channel needs to be made prior to assignment
}

func murderOccupant(n node) {
	n.channelToResident <- "BANG!"
}

func (n node) String() string {
	emptyNodeSymbol = "-"
	if n.resident == nil {
		return emptyNodeSymbol
	}
	return n.resident
}
