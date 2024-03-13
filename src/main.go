package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func dividir ( numbers []int, partes int){
    
    var slices [][]int
    var wgInner sync.WaitGroup
    
    for i := 0; i < partes; i++ {
        start := i * len(numbers) / partes
        end := (i + 1) * len(numbers) / partes
        newSlice := make([]int, len(numbers[start:end]))
        copy(newSlice, numbers[start:end])
        slices = append(slices, newSlice)
    }

    fmt.Println(slices)

    wgInner.Add(len(slices))

    for _, newSlice := range slices {
        go contar(newSlice, &wgInner)
    }

    // Espera a que todas las goroutines interiores terminen
    wgInner.Wait()
}

func contar(numbers []int, wg *sync.WaitGroup) {
    defer wg.Done() // Se asegura de que se llame a Done cuando la función retorne

    var value int
    for _, sum := range numbers {
        value += sum
    }
    fmt.Printf("El valor es: %d\n", value)
}

func main() {
    var numbers []int
    for i := 0 ; i < 100; i++{
        numbers = append(numbers, rand.Intn(10))
    }  
    fmt.Println(numbers)

    dividir(numbers,5)
}

// Otra alternativa es 

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func sumPartial(arr []int, start, end int, result chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := start; i < end; i++ {
// 		sum += arr[i]
// 	}
// 	result <- sum
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// numWorkers := 4
	// result := make(chan int, numWorkers)
// 	var wg sync.WaitGroup
// 	wg.Add(numWorkers)

// 	size := len(arr) / numWorkers
// 	for i := 0; i < numWorkers; i++ {
// 		start := i * size
// 		end := (i + 1) * size
// 		if i == numWorkers-1 {
// 			end = len(arr)
// 		}
// 		go sumPartial(arr, start, end, result, &wg)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()

// 	total := 0
// 	for partialSum := range result {
// 		total += partialSum
// 	}
// 	fmt.Println("Total sum:", total)
// }