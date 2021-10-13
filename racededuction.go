package main

import "sync"

func main() {
	count := 0
	var w sync.WaitGroup
	//var mux sync.Mutex
	w.Add(2)
	go func() {
		//mux.Lock()
		count += 1
		//mux.Unlock()
		w.Done()
	}()

	go func() {
		//mux.Lock()
		count += 1
		//mux.Unlock()
		w.Done()
	}()
	w.Wait()
	println(count)
}
