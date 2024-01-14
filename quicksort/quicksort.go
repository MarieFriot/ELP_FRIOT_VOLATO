package quicksort

import (
	"sync"
)

// Fonction permettant de spéparer une liste d'entier en un nombre donnée (listNumber) de
// soous-liste. La focntion renvoi un tableau où chaques lignes représentent une sous-liste.
func splitList(liste []int, listNumber int) [][]int {
	//Créer le tableau de résultat
	res := make([][]int, listNumber)

	//Si le nombre de liste demandé est 1 : il n'y a rien à faire.
	if listNumber == 1 {
		res[0] = liste[:]

	} else {
		//Calcul des indices ou la liste va être séparée
		//On créer une liste (separator) contenant chacun de ces indices.
		separator := []int{}
		for i := 1; i < listNumber; i++ {
			separator = append(separator, len(liste)/listNumber*i)
		}

		// On remplit le tableau de résultat avec les sous liste
		res[0] = liste[:separator[0]] // de 0 à separator[0] exclu
		for i := 1; i < listNumber-1; i++ {
			res[i] = liste[separator[i-1]:separator[i]]
		}
		res[listNumber-1] = liste[separator[listNumber-2]:]
	}
	return res
}

// Fonction permettant de trier une liste à l'aide de l'algorithme du quicksort
// Prend en parametre un wait groupe pour être exécuter plusieurs fois en parallele
func Quicksort(liste []int, wg *sync.WaitGroup) {

	defer wg.Done() // (wg externe décrementé qu'une fois car les appels récursifs se font avec un autre)

	// Si la liste est de taille 1 : on arrête
	if len(liste) <= 1 {
		return
	}

	pivot := liste[0]
	var liste1, liste2 []int

	// On partitionne les valeur de la liste dans deux liste :
	// la liste1 contient les valeurs plus petite que le pivot
	// la liste 2 contient les valeurs supérieurs ou égales à celle du pivot.
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

	//On ajoute a la liste1 le pivot
	liste1 = append(liste1, pivot)
	// On ajoute à la liste1 les valeurs de la liste2
	liste1 = append(liste1, liste2...) //... décompose les elements individuels de liste2 et les ajoute à liste1
	// liste prend les valeurs de liste1
	copy(liste, liste1)

}

func QuicksortParallel(liste []int, listNumber int) []int {
	var wgPartition sync.WaitGroup
	var wgLecture sync.WaitGroup

	// Si la taille de la liste est plus petit au égale à 1 on a rien à faire
	if len(liste) <= 1 {
		return liste
	}

	// On récupère les sous-liste (il y en a listNumber)
	tab := splitList(liste, listNumber)

	// On créer des canneaux pour envoyer les listes résultats des partions
	lowCh := make(chan []int, 2)
	upCh := make(chan []int, 2)

	// lLow va contenir toutes les valeurs de la liste à trier inférieurs au pivot
	lLow := []int{}
	//lUp va contenir toutes les valeurs de la liste à trier supérieurs au pivot
	lUp := []int{}

	// On créer une go routine pour lire les listes résultats envoyées dans les canneaux
	wgLecture.Add(1)
	go func() {
		for i := 0; i < listNumber; i++ {
			l1 := <-lowCh
			l2 := <-upCh
			lLow = append(lLow, l1...)
			lUp = append(lUp, l2...)
		}
		defer wgLecture.Done()
	}()

	// On choisit le pivot autour duquel toutes les listes vont être partitionnées
	pivot := liste[0]

	// On créer un nombre de go routine correspondant au nombre de sous-liste pour partitionner
	// chaques sous liste au tour de la valeur pivot
	for i := 0; i < listNumber; i++ {
		wgPartition.Add(1)
		go Partition(tab[i], pivot, &wgPartition, lowCh, upCh)
	}

	// On attend que toutes les go routines aient partitionner les sous-liste
	wgPartition.Wait()
	// On attent que les résultats de toutes les partitions aient été récupérés dans les canneaux
	wgLecture.Wait()

	// On effectue en parallele le Quicksort sur la liste contenant toutes les valeurs inférieurs au pivot
	// et la liste contenant toutes les valeurs supérieures au pivot.
	var wgQuicksort sync.WaitGroup
	wgQuicksort.Add(2)
	go Quicksort(lLow, &wgQuicksort)
	go Quicksort(lUp, &wgQuicksort)
	wgQuicksort.Wait()

	lLow = append(lLow, lUp...)
	return lLow

}

// version sequentielle du Quicksort
func QuicksortSeq(liste []int) {

	// Si la liste est de taille 0 ou 1 on a rien à faire
	if len(liste) <= 1 {
		return
	}

	pivot := liste[0]
	var liste1, liste2 []int

	// On partitionne les valeur de la liste dans deux liste :
	// la liste1 contient les valeurs plus petite que le pivot
	// la liste 2 contient les valeurs supérieurs ou égales à celle du pivot.
	for i := 1; i < len(liste); i++ {
		if liste[i] < pivot {
			liste1 = append(liste1, liste[i])
		} else {
			liste2 = append(liste2, liste[i])
		}
	}

	// Appel récursif sur les listes partionnées
	QuicksortSeq(liste1)
	QuicksortSeq(liste2)

	//On assemble les résultats
	liste1 = append(liste1, pivot)
	liste1 = append(liste1, liste2...) //... décompose les elements individuels de liste2 et les ajoute à liste1
	copy(liste, liste1)

}

// Fonction permettant de partitioner une liste en deuc listes au tour d'une valeur pivot
// la première liste contient les valeurs plus petite que le pivot
// la seconde liste  contient les valeurs supérieurs ou égales à celle du pivot.
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
	// On envoi dans le canal chLow la liste contenant les valeurs plus petite que le pivot
	chLow <- liste1
	// On envoi dans le canal chUp la liste contenant les valeurs plus petite que le pivot
	chUp <- liste2

}
