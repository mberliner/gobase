package main

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

var ch = make(chan byte)
var counter int64 = 0

func cuentaYBloquea() {
	atomic.AddInt64(&counter, 1)
	<-ch // Block this goroutine
}

func main() {

	var nroGoRutine int64 = 100000

	// Limit the number of spare OS threads to just 1
	fmt.Println(runtime.GOMAXPROCS(1))

	// Make a copy of MemStats
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)

	t0 := time.Now().UnixNano()
	for i := 0; i < int(nroGoRutine); i++ {
		go cuentaYBloquea()
	}
	runtime.Gosched()
	t1 := time.Now().UnixNano()
	runtime.GC()

	// Make a copy of MemStats
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	fmt.Println("\nNro of Goroutines", runtime.NumGoroutine())
	if atomic.LoadInt64(&counter) != nroGoRutine {
		fmt.Fprintf(os.Stderr, "failed to begin execution of all goroutines")
		os.Exit(1)
	}

	fmt.Printf("Per goroutine:\n")
	fmt.Printf("  Memory: %.2f bytes\n", float64(m1.Sys-m0.Sys)/float64(counter))
	fmt.Printf("  Time:   %f µs\n", float64(t1-t0)/float64(counter)/1e3)
}
