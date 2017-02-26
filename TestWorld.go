package main

func main() {
	w := createDefaultWorld()
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(w)
	}
}
