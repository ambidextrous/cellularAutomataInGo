package main

type species struct {
	name     string
	lifeTime int
	fitness  float32
	node     node
}

//func setLifetime() int {
//	maxLifetime := 10000 //Recommended value: 10000
//	return rand.Intn(maxLifetime)
//}
//
//func setName() {
//	name := "1"
//	return name
//}
//
//func setFitness() {
//	fitness := 0.8 // Recommended value: 0.8
//	return fitness
//}

func spawn(s species, n node) {
	return species{name: s.name, lifeTime: s.lifeTime, fitness: s.fitness, node: n}
}

func (s species) String() string {
	return s.name
}
