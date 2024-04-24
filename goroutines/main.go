package main

import (
	"fmt"
	"sync"
	"time"
)

// this entire program is a simulation of a database with some data and queries

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var db = []string{"id0", "id1", "id2", "id3", "id4"}
var results []string // this will store the results of the queries

func query(id int) {
	// simulate some delay in querying the database
	var delay int32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(db[id])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

// this function will log the current results
// we use a read lock because we don't want to modify the results while we are reading them
// this is because the results are being modified by the save function concurrently
// if we don't use a lock, we might get inconsistent results (a race condition)
func log() {
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}

func main() {
	t0 := time.Now()
	for i := 0; i < len(db); i++ {
		// query(i) runs sequentially (one after the other)
		wg.Add(1)   // this will wait for the query to finish
		go query(i) // this will run the query concurrently, but the main function will not wait for the result
	}
	wg.Wait()
	fmt.Printf("\nTime taken: %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", results)
	fmt.Println(len(results))
}
