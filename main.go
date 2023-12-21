package main 

import (
  "fmt"
  "ELP_FRIOT_VOLATO/quicksort" //notre bibliotheque avec la fonction quicksort dedans
  "sync"
)

//Parallel quicksort version 2 processeur
func ParallelQSv2(data []int, wg *sync.WaitGroup, channel chan []int) {
  defer wg.Done()

  if len(data) <= 1 {
    ch <- data
    return
  }

  var low, high []int
  low, high := quicksort.Partition(data)

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

  result := append(<- lowerCh, pivot)
  result = append(result, <- upperCh)

  ch <- result
  }
  
}
