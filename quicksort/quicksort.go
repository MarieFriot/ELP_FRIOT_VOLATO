package quicksort

import (
	"sync"
)

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

	var wgInner sync.WaitGroup
	wgInner.Add(2)
	go Quicksort(liste1, &wgInner)
	go Quicksort(liste2, &wgInner)
	wgInner.Wait()

	liste1 = append(liste1, pivot)
	liste1 = append(liste1, liste2...) //... décompose les elements individuels de liste2 et les ajoute à liste1
	copy(liste, liste1)

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

// coupe une liste en deux avec un liste d'element inferieur au pivot et une superieur au pivot, les deux listes ne sont pas range.
func Partition(liste []int, pivot int, wg *sync.WaitGroup) ([]int, []int) {
	defer wg.Done()
	var liste1, liste2 []int
	for i := 0; i < len(liste); i++ {
		if liste[i] <= pivot {
			liste1 = append(liste1, liste[i])
		} else {
			liste2 = append(liste2, liste[i])
		}
	}
	return liste1, liste2
}
