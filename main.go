package main

import (
	"fmt"
	"math/rand"
	"sync"
	"test_quicksort_parav2/quicksort"
	"time"
)

func splitList(liste []int, listNumber int) [][]int {
	res := make([][]int, listNumber)
	separator := []int{}
	for i := 1; i < listNumber; i++ {
		separator = append(separator, len(liste)/listNumber*i)
	}
	fmt.Println(separator)
	res[0] = liste[:separator[0]]
	for i := 1; i < listNumber-1; i++ {
		res[i] = liste[separator[i-1]:separator[i]]
	}
	res[listNumber-1] = liste[separator[listNumber-2]:]
	return res
}

func bigListeGen(size int) []int {
	liste := []int{}
	for i := 0; i < size; i++ {
		liste = append(liste, rand.Intn(1000))
	}
	return liste
}

func QuicksortParallel(liste []int, listNumber int) []int {
	var wg sync.WaitGroup

	if len(liste) <= 1 {
		return liste
	}
	tab := splitList(liste, listNumber)
	pivot := liste[0]

	lowCh := make(chan []int, 2)
	upCh := make(chan []int, 2)

	lLow := []int{}
	lUp := []int{}

	go func() {
		for i := 0; i < listNumber; i++ {
			l1 := <-lowCh
			l2 := <-upCh
			lLow = append(lLow, l1...)
			lUp = append(lUp, l2...)
			fmt.Println("Received low:", l1)
			fmt.Println("Received up:", l2)
		}
	}()

	wg.Add(listNumber)
	for i := 0; i < listNumber; i++ {
		go quicksort.Partition(tab[i], pivot, &wg, lowCh, upCh)
	}

	wg.Wait()

	wg.Add(2)
	go quicksort.Quicksort(lLow, &wg)
	go quicksort.Quicksort(lUp, &wg)
	wg.Wait()

	lLow = append(lLow, lUp...)
	return lLow

}

func main() {
	liste3 := bigListeGen(10)
	liste4 := make([]int, len(liste3))
	copy(liste4, liste3)
	fmt.Println(liste3)

	startTime := time.Now()
	result := QuicksortParallel(liste3, 4)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Temps d'exécution parralllele : %v\n", duration)
	fmt.Println(result)
	fmt.Printf("%v\n", result[0])

	startTime = time.Now()
	quicksort.QuicksortSeq(liste4)
	endTime = time.Now()
	duration = endTime.Sub(startTime)
	fmt.Printf("Temps d'exécution sequentiel : %v\n", duration)
	fmt.Println(liste4)

}
