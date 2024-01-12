package quicksort

import (
	"sync"
)

func splitList(liste []int) ([]int, []int) {
	milieu := len(liste) / 2
	return liste[:milieu], liste[milieu:]
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


	///On peut également faire les aapppeles récursifs sans créer de go routines pour gagner du temps !
	var wgInner sync.WaitGroup
	wgInner.Add(2)
	go Quicksort(liste1, &wgInner)
	go Quicksort(liste2, &wgInner)
	wgInner.Wait()

	liste1 = append(liste1, pivot)
	liste1 = append(liste1, liste2...) //... décompose les elements individuels de liste2 et les ajoute à liste1
	copy(liste, liste1)

}

func QuicksortParallel(liste []int) []int {
	var wg sync.WaitGroup
	var wgCom sync.WaitGroup

	if len(liste) <= 1 {
		return liste
	}

	lLow := []int{}
	lUp := []int{}

	lowCh := make(chan []int, 2)
	upCh := make(chan []int, 2)

	wgCom.Add(1)
	go func() {
		for i := 0; i < 2; i++ {
			l1 := <-lowCh
			l2 := <-upCh
			lLow = append(lLow, l1...)
			lUp = append(lUp, l2...)
			//fmt.Println("Received low:", l1)
			//fmt.Println("Received up:", l2)
		}
		defer wgCom.Done()
	}()

	l1, l2 := splitList(liste)
	pivot := liste[0]

	wg.Add(2)
	go Partition(l1, pivot, &wg, lowCh, upCh)
	go Partition(l2, pivot, &wg, lowCh, upCh)
	wg.Wait()
	wgCom.Wait()

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
	print("start")
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
	print("done")

}
