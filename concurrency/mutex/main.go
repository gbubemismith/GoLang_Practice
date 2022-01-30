package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Mutex -> Mutual exclusion, it provides a way to lock and unlock data
func main() {
  rand.Seed(time.Now().UnixNano())

  var wg sync.WaitGroup

  counter := 0
  for i := 0; i < 5; i++ {
	  wg.Add(1)
	  counter += 1
	  go func ()  {
		  defer func() {
			  fmt.Println()
		  }
	  }
  }
}
