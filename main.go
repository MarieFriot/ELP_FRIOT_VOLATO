package main

import (
	"ELP_FRIOT_VOLATO-Golang/quicksort" //notre bibliotheque avec la fonction quicksort dedans
	"fmt"
	"math/rand"
	"sync"
)

func bigListeGen() []int {
	liste := []int{}
	for i := 0; i < 10; i++ {
		liste = append(liste, rand.Intn(100))
	}
	return liste
}
func splitList(liste []int)([]int, []int) {
  milieu int
  res1 int[]
  res2 int[]
  milieu = len(liste)/2
  for i=0; i< milieu; i++{
    res1.append(liste[i])
  }
  for i = milieu; i< len(liste); i++{
    res2.append(liste[i])
  }
  return res1,res2
}

// Parallel quicksort version 2 processeur
func ParallelQSv2(data []int) {
  //séparer liste en deux
  l1, l2 := splitList(data)

  channel := make(chan int[])
  var waitGroup sync.WaitGroup
  waitGroup.Add(2)
  
  //premiere go routine
  go func()  {
    defer waitGroup.Done()
    l1Low, l1Up :=partition(l1)
    channel <- l1Up //envoie de la plus grande liste dans le channel
    l2Low := <-channel //recoit la plus petite liste 
    
  }()

  //deuxieme go routine
  go func() {
    defer waitGroup.Done()
    l2Low, l2Up :=partition(l2)
    channel <- l2Low //envoie de la plus petite liste dans le channel
    l1Up := <-channel //recoie la plus grande liste 
  }()
  
  waitGroup.Wait()
}



func main() {
	liste := bigListeGen()
	
	fmt.Println("liste_trie", liste_trie)
}
