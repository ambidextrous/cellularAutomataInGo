package main

World struct {
	[][]Node node
	int height
	int width
	[][]int oneSpaceMoves
}

func populate(w world) {

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



