package main 

import (
  "fmt"
  "ELP_FRIOT_VOLATO/quicksort" //notre bibliotheque avec la fonction quicksort dedans
  "sync"
)

func main() {
  liste := []int{1,5,9,8,3,6,4,7,2}
  quicksort.Quicksort(liste)
  fmt.Println(liste)
}
