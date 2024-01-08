package main

import (
	"fmt"
	"math/rand"
	"test_quicksort_parav2/quicksort"
	"time"
)

func bigListeGen(size int) []int {
	liste := []int{}
	for i := 0; i < size; i++ {
		liste = append(liste, rand.Intn(1000))
	}
	return liste
}

func main() {
	liste3 := bigListeGen(10000)
	liste4 := make([]int, len(liste3))
	copy(liste4, liste3)
	//fmt.Println(liste3)

	startTime := time.Now()
	result := quicksort.QuicksortParallel(liste3, 10)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Temps d'exécution parralllele : %v\n", duration)
	//fmt.Println(result)
	fmt.Printf("%v\n", result[0])

	startTime = time.Now()
	quicksort.QuicksortSeq(liste4)
	endTime = time.Now()
	duration = endTime.Sub(startTime)
	fmt.Printf("Temps d'exécution sequentiel : %v\n", duration)
	//fmt.Println(liste4)

}
