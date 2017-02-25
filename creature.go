package main

import "fmt"
import "time"
import "math/rand" 

//type species interface {
//	area() float64
//	setLifetime() int
//	setName() string
//	setFitness() float64
//	spawn() species
//}

// Abstract form unnessessary: instanciate only as Species1, etc.
//type Creature struct {
//	string  name
//	int     lifeTime
//	float32 fitness
//	Node    node
//}


type species struct {
	string  name
	int     lifeTime
	float32 fitness
	node    node
	<-chan string channelFromNode
}

func run(s species) {
	start := time.Now()
	for {
		select {
			// If lifetime over, reproduce and return
			case time.Now() > start + s.lifeTime:
				fmt.Println("A child is born!")
				reproduce(s)
				return 
			// If interrupted, return without reproducing
			case <-s.channelFromNode:
				fmt.Println("Alas, for I am slain!")
				s.node.resident = nil
				return
			// Otherwise, do nothing
			default: // Prevents goroutine from blocking
		}
	}
}

func populateNode(parent species, n node) {
	child := spawn(parent,n)
	setResident(n,child)
}


func spawn(s species, n node) {
	return species{name: s.name, lifeTime: s.lifeTime, fitness: s.fitness, node: n}
}

func reproduceInIndividualNodes(n node, s species) {
	for neighbouringNode := range getNeighbouringNodes(n) {
		if isEmpty(neighbouringNode) {
			if rand.float64() <= s.fitness {
				populateNode(s, neighbouringNode)
			} 
		}
		else {
			other := getResident(neighbouringNode)
			if rand.float64() <= c.fitness - other.fitness {
				murderOccupant(neighbouringNode) 
				populateNode(s, neighbouringNode)
			}  
		}
	}
}

func reproduce(s species) {
	setEmpty(s.node)
	reproduceInIndividualNodes(s)
}

func (s species) String() string {
	return s.name
}

