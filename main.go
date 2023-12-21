package main 

import (
  "fmt"
  "ELP_FRIOT_VOLATO/quicksort" //notre bibliotheque avec la fonction quicksort dedans
  "sync"
  "math/rand"
)

func bigListeGen() []int {
  liste := []int
  for i:= 0; i< 100; i++{
    liste = append(liste, rand.Intn(100))
    }
  return liste
  }
//Parallel quicksort version 2 processeur
func ParallelQSv2(data []int, wg *sync.WaitGroup, channel chan []int) {
  defer wg.Done()

  if len(data) <= 1 {
    channel <- data
    return
  }

  var low, high []int
  low, high = quicksort.Partition(data)

  lowerWg := &sync.WaitGroup{}
  upperWg := &sync.WaitGroup{}
  lowerCh := make(chan []int)
  upperCh := make(chan []int)

  lowerWg.Add(1)
  go ParallelQSv2(low, lowerWg, lowerCh)

  upperWg.Add(1)
  go ParallelQSv2(high, upperWg, upperCh)

  lowerWg.Wait()
  upperWg.Wait()

  result := append(<- lowerCh, nil)
  result = append(result, <- upperCh...)

  channel <- result
  }
  
}
func main{
  liste := bigListeGen()
  var waitGroup sync.WaitGroup
  channel := make(chan []int)
  waitGroup.Add(1)
  go ParallelQSv2(liste, &waitGroup, channel)
  go func() {
    waitGroup.Wait()
    close(channel)
  }()
  liste_trie := <-channel
  fmt.Println("liste_trie", liste_trie)
  }
  
  
