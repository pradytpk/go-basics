package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count += 1
	fmt.Println("Served iteration", iteration)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	hitsCounter := Hits{}
	for i := 0; i < 20; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitsCounter, iteration)
	}
	fmt.Printf("waiting for goroutines..\n\n")
	wg.Wait()

	hitsCounter.Lock()
	totalHits := hitsCounter.count
	hitsCounter.Unlock()
	fmt.Println("Total hits", totalHits)
}
