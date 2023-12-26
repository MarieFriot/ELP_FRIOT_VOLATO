package main

import (
	"fmt"
	"math/rand"
	"sync"
	"ELP_FRIOT_VOLATO-Golang/quicksort"
	"time"
)

func splitList(liste []int) ([]int, []int) {
	milieu := len(liste) / 2
	return liste[:milieu], liste[milieu:]
}

func bigListeGen(size int) []int {
	liste := []int{}
	for i := 0; i < size; i++ {
		liste = append(liste, rand.Intn(1000))
	}
	return liste
}

func QuicksortParallel(liste []int) []int {
	var wg sync.WaitGroup

	if len(liste) <= 1 {
		return liste
	}
	l1, l2 := splitList(liste)
	pivot := liste[0]

	l1Ch := make(chan struct{ l1Low, l1Up []int })
	l2Ch := make(chan struct{ l2Low, l2Up []int })

	wg.Add(2)
	go func() {
		l1Low, l1Up := quicksort.Partition(l1, pivot, &wg)
		l1Ch <- struct{ l1Low, l1Up []int }{l1Low, l1Up}
	}()
	go func() {
		l2Low, l2Up := quicksort.Partition(l2, pivot, &wg)
		l2Ch <- struct{ l2Low, l2Up []int }{l2Low, l2Up}
	}()
	wg.Wait()

	result1 := <-l1Ch
	result2 := <-l2Ch

	lLow := append(result1.l1Low, result2.l2Low...)
	lUp := append(result1.l1Up, result2.l2Up...)

	wg.Add(2)
	go quicksort.Quicksort(lLow, &wg)
	go quicksort.Quicksort(lUp, &wg)
	wg.Wait()

	lLow = append(lLow, lUp...)
	return lLow

}

func main() {
	liste3 := bigListeGen(100000)
	liste4 := make([]int, len(liste3))
	copy(liste4, liste3)

	startTime := time.Now()
	result := QuicksortParallel(liste3)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Temps d'exécution parralllele : %v\n", duration)
	fmt.Println(result[0])

	startTime = time.Now()
	quicksort.QuicksortSeq(liste4)
	endTime = time.Now()
	duration = endTime.Sub(startTime)
	fmt.Printf("Temps d'exécution sequentiel : %v\n", duration)

}
