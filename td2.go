package main

import (
    "fmt"
	"math/rand"
	"sync"
)

//2
// func run(c chan string, name string) {
// 	c <- name
// }

// func main() {
// 	channel := make(chan string)	
// 	go run(channel, "Bonjour, Goroutine !")
// 	fmt.Println(<-channel)
// }

//3
//if premier channel -> que écriture
// if dernier channel -> que lecture
// tous les autres lecture et ecriture
// recuperer le numéro du channel envoyé en paramètre
// var wg sync.WaitGroup

// func run(id, N int, c []chan int) {
// 	defer wg.Done()
// 	switch id {
// 	case 0:
// 		fmt.Println("Je suis le premier channel")
// 		c[0] <- id
// 	case N - 1:
// 		fmt.Println("Je suis le dernier channel : " + string (id+<-c[id-1]))
// 	default:
// 		fmt.Printf("Je suis le channel %d \n", id)
// 		v:= <-c[id-1]
// 		c[id] <- v + id
// 	}
// }

// func main() {
// 	N:= 10
// 	channel := make([]chan int, N)
// 	for i := 0; i < N; i++ {
// 		wg.Add(1)
// 		channel[i] = make(chan int, 1)// 1 pour buffer = 1 seul thread à la fois
// 		go run(i, N, channel)
// 	}
// 	wg.Wait()
// 	fmt.Println("Fini")
	
// }

//4
var wg sync.WaitGroup

func maitre(id, N int, c []chan int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	c[id]<- randomNumber
	fmt.Printf("Maitre %d envoie %d\n", id, randomNumber)

}

func ouvrier(id, N int, c []chan int) {
	defer wg.Done()

}

func main() {
	M:= 10
	channelEnvoie := make([]chan int, M)
	for i := 0; i < M; i++ {
		wg.Add(1)
		channelEnvoie[i] = make(chan int, 1)// 1 pour buffer = 1 seul thread à la fois
		go run(i, N, channelEnvoie)
	}
	wg.Wait()
	fmt.Println("Fini")
	
}
