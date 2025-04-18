package main

import (
    "fmt"
	//"math/rand"
	"sync"
	//"time"
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
// var wg sync.WaitGroup

// // Fonction pour les ouvrières
// func ouvrier(id int, in chan int, out chan int) {
//     defer wg.Done()
//     // Recevoir le nombre du maître
//     num := <-in
//     fmt.Printf("Ouvrier %d a reçu %d\n", id, num)
//     // Calculer le carré et l'envoyer au maître
//     out <- num * num
// }

// // Fonction principale
// func main() {
//     rand.Seed(time.Now().UnixNano()) // Initialiser la graine pour les nombres aléatoires

//     M := 5 // Nombre d'ouvriers
//     inChannels := make([]chan int, M)  // Channels pour envoyer les nombres aux ouvriers
//     outChannels := make([]chan int, M) // Channels pour recevoir les résultats des ouvriers

//     // Initialiser les channels et démarrer les ouvriers
//     for i := 0; i < M; i++ {
//         inChannels[i] = make(chan int, 1)
//         outChannels[i] = make(chan int, 1)
//         wg.Add(1)
//         go ouvrier(i, inChannels[i], outChannels[i])
//     }

//     // Goroutine maître
//     go func() {
//         for i := 0; i < M; i++ {
//             randomNumber := rand.Intn(100) // Générer un nombre aléatoire
//             fmt.Printf("Maître envoie %d à l'ouvrier %d\n", randomNumber, i)
//             inChannels[i] <- randomNumber // Envoyer le nombre à l'ouvrier
//         }
//     }()

//     // Recevoir les résultats des ouvriers
//     for i := 0; i < M; i++ {
//         result := <-outChannels[i]
//         fmt.Printf("Maître a reçu %d de l'ouvrier %d\n", result, i)
//     }

//     // Attendre la fin des ouvriers
//     wg.Wait()
//     fmt.Println("Fini")
// }

//5.1 
// func goroutine(id, P, K int, ch chan int, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for i := 0; i < K; i++ {
//         msg := <-ch // Recevoir le message
//         fmt.Printf("Goroutine %d a reçu %d (itération %d)\n", id, msg, i+1)
//         msg++ // Incrémenter le message
//         ch <- msg // Envoyer au suivant
//     }
// }

// func main() {
//     const P = 5 // Nombre de goroutines
//     const K = 3 // Nombre de tours du message

//     ch := make(chan int, 1) // Channel partagé
//     var wg sync.WaitGroup

//     // Créer les goroutines
//     for i := 0; i < P; i++ {
//         wg.Add(1)
//         go goroutine(i, P, K, ch, &wg)
//     }

//     // Démarrer le message
//     ch <- 0 // Envoyer le message initial

//     // Attendre la fin des goroutines
//     wg.Wait()
//     fmt.Println("Fini")
// }

// //5.2 
// func goroutine(id, P, K int, in, out chan int, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for i := 0; i < K; i++ {
//         msg := <-in // Recevoir le message
//         fmt.Printf("Goroutine %d a reçu %d (itération %d)\n", id, msg, i+1)
//         msg++ // Incrémenter le message
//         out <- msg // Envoyer au suivant
//     }
// }

// func main() {
//     const P = 5 // Nombre de goroutines
//     const K = 3 // Nombre de tours du message

//     channels := make([]chan int, P) // Un channel par goroutine
//     var wg sync.WaitGroup

//     // Initialiser les channels
//     for i := 0; i < P; i++ {
//         channels[i] = make(chan int, 1)
//     }

//     // Créer les goroutines
//     for i := 0; i < P; i++ {
//         wg.Add(1)
//         go goroutine(i, P, K, channels[i], channels[(i+1)%P], &wg)
//     }

//     // Démarrer le message
//     channels[0] <- 0 // Envoyer le message initial

//     // Attendre la fin des goroutines
//     wg.Wait()
//     fmt.Println("Fini")
// }

//6
// Fonction pour fusionner deux tableaux triés
func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    // Ajouter les éléments restants
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}

// Fonction mergeSort parallèle
func mergeSort(arr []int, wg *sync.WaitGroup, ch chan []int) {
    defer wg.Done()

    // Si le tableau contient un seul élément ou est vide, il est déjà trié
    if len(arr) <= 1 {
        ch <- arr
        return
    }

    // Diviser le tableau en deux parties
    mid := len(arr) / 2
    leftChan := make(chan []int, 1)
    rightChan := make(chan []int, 1)

    var subWg sync.WaitGroup
    subWg.Add(2)

    // Lancer les goroutines pour trier chaque moitié
    go mergeSort(arr[:mid], &subWg, leftChan)
    go mergeSort(arr[mid:], &subWg, rightChan)

    // Attendre que les deux moitiés soient triées
    subWg.Wait()
    close(leftChan)
    close(rightChan)

    // Fusionner les résultats
    left := <-leftChan
    right := <-rightChan
    ch <- merge(left, right)
}

func main() {
    // Tableau à trier
    arr := []int{38, 27, 43, 3, 9, 82, 10}

    fmt.Println("Tableau initial :", arr)

    // Channel pour récupérer le tableau trié
    ch := make(chan []int, 1)
    var wg sync.WaitGroup

    // Lancer le tri fusion parallèle
    wg.Add(1)
    go mergeSort(arr, &wg, ch)

    // Attendre la fin du tri
    wg.Wait()
    close(ch)

    // Récupérer le tableau trié
    sortedArr := <-ch
    fmt.Println("Tableau trié :", sortedArr)
}