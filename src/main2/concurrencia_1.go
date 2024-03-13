package main

import (
	"fmt"
	"time"
)


func canalFuncion(c chan<- string, nombre string) {
	c <- nombre
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	out := make(chan string)

	go canalFuncion(c1, "Pepito")
	go canalFuncion(c2, "Jose Fernando el mamahuevo numero 1")

	go func() {
		for {
			select {

			case v := <-c1:
				time.Sleep(4 * time.Second) //Los "case" no son en paralelo, una vez que recibe al primero que se ejecuta se ejecuta secuencial, si llegase a terminar otra goroutine debera esperar a que el case termine.
				out <- v
			case v := <-c2:
				out <- v

			}
		}
	}()

	for i := 0; i < 2; i++ {
		fmt.Println(<- out)
	}
}