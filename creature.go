package main

import "fmt"
import "time"

type species interface {
	area() float64
	setLifetime() int
	setName() string
	setFitness() float64
	spawn() species
}

// Abstract form unnessessary: instanciate only as Species1, etc.
type Creature struct {
	string  name
	int     lifeTime
	float32 fitness
	Node    node
}

func getName(s species) string {
	return s.name
} 

func run(s species) {
	time.Sleep(s.lifeTime)	
	reproduce(s)
}

func populateNode(n node) {
	s := spawn(n)
}

func reproduceInIndividualNodes(n) {
	for neighbouringNode := range getNeighbouringNodes(n) {
		if isEmpty(neighbouringNode) {
			if 
		}
	}
	
}

func getFitness(s species) float32 {
	return s.fitness
}

func reproduce(s species) {
	setEmpty(s.node)
	reproduceInIndividualNodes(s)
}
