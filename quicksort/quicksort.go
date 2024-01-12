package quicksort

import (
	"sync"
)

func splitList(liste []int, listNumber int) [][]int {
	res := make([][]int, listNumber)
	if listNumber == 1 {
		res[0] = liste[:]
	} else {
		separator := []int{}
		for i := 1; i < listNumber; i++ {
			separator = append(separator, len(liste)/listNumber*i)
		}
		res[0] = liste[:separator[0]]
		for i := 1; i < listNumber-1; i++ {
			res[i] = liste[separator[i-1]:separator[i]]
		}
		res[listNumber-1] = liste[separator[listNumber-2]:]
	}
	return res
}

func Quicksort(liste []int, wg *sync.WaitGroup) {
	defer wg.Done() //wg externe décrementé qu'une fois car les appels récursifs se font avec un autre
	if len(liste) <= 1 {
		return
	}
	pivot := liste[0]

	var liste1, liste2 []int

	for i := 1; i < len(liste); i++ {
		if liste[i] < pivot {
			liste1 = append(liste1, liste[i])
		} else {
			liste2 = append(liste2, liste[i])
		}
	}

	//Ou faire des appels récursifs de Quicksort avec un waitgroup
	//var wgInner sync.WaitGroup
	//wgInner.Add(2)
	QuicksortSeq(liste1)
	QuicksortSeq(liste2)
	//wgInner.Wait()

	liste1 = append(liste1, pivot)
	liste1 = append(liste1, liste2...) //... décompose les elements individuels de liste2 et les ajoute à liste1
	copy(liste, liste1)

}

func QuicksortParallel(liste []int, listNumber int) []int {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	if len(liste) <= 1 {
		return liste
	}
	print("hi")
	tab := splitList(liste, listNumber)
	pivot := liste[0]

	lowCh := make(chan []int, 2)
	upCh := make(chan []int, 2)

	lLow := []int{}
	lUp := []int{}

	wg2.Add(1)
	go func() {
		for i := 0; i < listNumber; i++ {
			l1 := <-lowCh
			l2 := <-upCh
			lLow = append(lLow, l1...)
			lUp = append(lUp, l2...)
		}
		defer wg2.Done()
	}()

	for i := 0; i < listNumber; i++ {
		wg.Add(1)
		go Partition(tab[i], pivot, &wg, lowCh, upCh)
	}

	wg.Wait()
	wg2.Wait()
	wg.Add(2)
	go Quicksort(lLow, &wg)
	go Quicksort(lUp, &wg)
	wg.Wait()

	lLow = append(lLow, lUp...)
	return lLow

}

// /version sequentielle
func QuicksortSeq(liste []int) {
	if len(liste) <= 1 {
		return
	}
	pivot := liste[0]

	var liste1, liste2 []int

	for i := 1; i < len(liste); i++ {
		if liste[i] < pivot {
			liste1 = append(liste1, liste[i])
		} else {
			liste2 = append(liste2, liste[i])
		}
	}

	QuicksortSeq(liste1)
	QuicksortSeq(liste2)
	liste1 = append(liste1, pivot)
	liste1 = append(liste1, liste2...) //... décompose les elements individuels de liste2 et les ajoute à liste1
	copy(liste, liste1)

}

func Partition(liste []int, pivot int, wg *sync.WaitGroup, chLow chan []int, chUp chan []int) {
	defer wg.Done()
	var liste1, liste2 []int
	for i := 0; i < len(liste); i++ {
		if liste[i] <= pivot {
			liste1 = append(liste1, liste[i])
		} else {
			liste2 = append(liste2, liste[i])
		}
	}
	chLow <- liste1
	chUp <- liste2

}
