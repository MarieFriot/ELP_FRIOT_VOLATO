package quicksort

import "fmt"

func main() {
	liste := []int{1, 3, 2, 4, 2, 5, 6}
	quicksort(liste)
	fmt.Println(liste)
}
func quicksort(liste []int) {
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
	quicksort(liste1)
	quicksort(liste2)

	liste1 = append(liste1, pivot)
	liste1 = append(liste1, liste2...) //... décompose les elements individuels de liste2 et les ajoute à liste1
	copy(liste, liste1)

}
